package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/forestgiant/oauthgrpc/ticktick/api"
)

var (
	tick1 = flag.String("tick1", "127.0.0.1", "Hostname/IP of the first tick server")
	tick2 = flag.String("tick2", "", "Hostname/IP of the optional second tick server")
	tick3 = flag.String("tick3", "", "Hostname/IP of the optional third tick server")
	tick4 = flag.String("tick4", "", "Hostname/IP of the optional fourth tick server")
	tick5 = flag.String("tick5", "", "Hostname/IP of the optional fifth tick server")
	port  = flag.Int("port", 9001, "The port to use when connecting to ticktick services")
)

func main() {
	flag.Parse()
	hosts := []string{*tick1}
	addHost := func(host string) {
		if len(host) > 0 {
			hosts = append(hosts, host)
		}
	}
	addHost(*tick2)
	addHost(*tick3)
	addHost(*tick4)
	addHost(*tick5)

	for _, host := range hosts {
		address := fmt.Sprintf("%s:%d", host, *port)
		ctx, cancelFunc := context.WithCancel(context.Background())
		log.Printf("Tick webserver connecting to: %s\n", address)
		tickTickClient, err := api.NewClient(ctx, address, nil)
		if err != nil {
			cancelFunc()
			log.Fatal(err)
		}
		defer cancelFunc()

		tickTickClient.GetMessages(func(msg string) {
			fmt.Println(msg)
		})
	}

	select {}
}
