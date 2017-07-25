package api

import (
	"context"
	"log"

	"io"

	"github.com/forestgiant/oauthgrpc/ticktick/pb"
	"google.golang.org/grpc"
)

type Client struct {
	rpc  pb.TickTickClient
	conn *grpc.ClientConn
}

func (c *Client) GetMessages(callback func(string)) error {
	ctx, cancelFunc := context.WithCancel(context.Background())
	stream, err := c.rpc.GetMessages(ctx, &pb.ReadMessagesRequest{})
	if err != nil {
		cancelFunc()
		return err
	}
	go func() {
		defer c.conn.Close()
		defer cancelFunc()
		for {
			rs, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				cancelFunc()
				log.Fatal(err)
			}
			if callback != nil {
				callback(rs.Message)
			}
		}
	}()

	return nil
}

func NewClient(ctx context.Context, ticktickAddress string, opts []grpc.DialOption) (*Client, error) {
	c := &Client{}
	if len(opts) == 0 {
		opts = append(opts, grpc.WithInsecure())
	}
	var err error
	c.conn, err = grpc.DialContext(ctx, ticktickAddress, opts...)
	if err != nil {
		return nil, err
	}
	c.rpc = pb.NewTickTickClient(c.conn)

	return c, nil
}
