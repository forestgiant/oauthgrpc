package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/thales17/oauthgrpc/ticktick/api"
)

var (
	address = flag.String("address", "127.0.0.1:9001", "Address of the server")
)

func main() {
	flag.Parse()
	ctx, cancelFunc := context.WithCancel(context.Background())
	tickTickClient, err := api.NewClient(ctx, *address, nil)
	if err != nil {
		cancelFunc()
		log.Fatal(err)
	}
	defer cancelFunc()
	tickTickClient.GetMessages(func(msg string) {
		fmt.Println(msg)
	})

	select {}
}
