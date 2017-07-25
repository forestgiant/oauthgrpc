package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/forestgiant/oauthgrpc/ticktick/pb"

	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 9001, "Foo server port")
	name = flag.String("name", "TickTickService", "The name of this instance of the ticktick service")
	tps  = flag.Int("tps", 1, "Ticks Per Second")
)

type tickTickClient struct {
	val chan int64
}

func newTickTickClient() *tickTickClient {
	c := &tickTickClient{}
	c.val = make(chan int64)

	return c
}

type tickTickServer struct {
	name     string
	duration time.Duration
	ticker   *time.Ticker
	val      int64
	clients  []*tickTickClient
	mutex    sync.Mutex
}

func (s *tickTickServer) GetMessages(req *pb.ReadMessagesRequest, stream pb.TickTick_GetMessagesServer) error {
	sendMessage := func(msg string, stream pb.TickTick_GetMessagesServer) error {
		message := &pb.TickTickMessage{}
		message.Message = msg
		if err := stream.Send(message); err != nil {
			return err
		}

		return nil
	}

	c := newTickTickClient()
	var val int64
	s.mutex.Lock()
	s.clients = append(s.clients, c)
	val = s.val
	s.mutex.Unlock()

	// Immediately let the stream know the current name and value
	if err := sendMessage(fmt.Sprintf("%s #%d", s.name, val), stream); err != nil {
		return err
	}
	ctx := stream.Context()

	for {
		select {
		case v := <-c.val:
			sendMessage(fmt.Sprintf("%s #%d", s.name, v), stream)
		case <-ctx.Done():
			s.mutex.Lock()
			// Remove this client from the list of clients stored in the server
			for i, sc := range s.clients {
				if sc == c {
					s.clients = append(s.clients[:i], s.clients[i+1:]...)
				}
			}
			s.mutex.Unlock()
			return nil
		}
	}
}

func newTickTickServer(name string, duration time.Duration) *tickTickServer {
	f := &tickTickServer{name: name, duration: duration}
	f.val = 0
	f.ticker = time.NewTicker(duration)
	// Start the ticker for the server
	go func() {
		for _ = range f.ticker.C {
			f.mutex.Lock()
			f.val++
			log.Printf("%s: %d\n", name, f.val)
			for _, c := range f.clients {
				c.val <- f.val
			}
			f.mutex.Unlock()
		}
	}()
	return f
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Starting ticktick services on port: %d with name: %s and ticks per second: %d\n", *port, *name, *tps)
	grpcServer := grpc.NewServer()
	pb.RegisterTickTickServer(grpcServer, newTickTickServer(*name, time.Second*time.Duration((*tps))))
	grpcServer.Serve(lis)
}
