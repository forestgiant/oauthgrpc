package main

import (
	"flag"
	"fmt"
	"log"

	"net"
	"sync"
	"time"

	"google.golang.org/grpc"

	"gitlab.fg/go/oauthgrpc/foo/pb"
)

var (
	port = flag.Int("port", 9001, "Foo server port")
)

type tickTickClient struct {
	val chan int
}

func newTickTickClient() *tickTickClient {
	c := &tickTickClient{}
	c.val = make(chan int)

	return c
}

type tickTickServer struct {
	name     string
	duration time.Duration
	ticker   *time.Ticker
	val      int
	clients  []*tickTickClient
	mutex    sync.Mutex
}

func (s *tickTickServer) GetMessages(req *pb.ReadMessagesRequest, stream pb.TickTick_GetMessagesServer) error {
	c := newTickTickClient()

	s.mutex.Lock()
	s.clients = append(s.clients, c)
	s.mutex.Unlock()
	ctx := stream.Context()
	for {
		select {
		case v := <-c.val:
			msg := &pb.TickTickMessage{}
			msg.Message = fmt.Sprintf("%s #%d", s.name, v)
			if err := stream.Send(msg); err != nil {
				return err
			}
		case <-ctx.Done():
			return nil
		}
	}
}

func newTickTickServer(name string, duration time.Duration) *tickTickServer {
	f := &tickTickServer{name: name, duration: duration}
	f.ticker = time.NewTicker(time.Second * 1)
	go func() {
		for _ = range f.ticker.C {
			f.val++
			f.mutex.Lock()
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
	grpcServer := grpc.NewServer()
	pb.RegisterTickTickServer(grpcServer, newTickTickServer("Foo", time.Second*1))
	grpcServer.Serve(lis)
}
