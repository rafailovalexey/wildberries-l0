package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nats-io/stan.go"
)

func main() {
	sc, err := stan.Connect("test-cluster", "subscriber-1", stan.NatsURL("http://localhost:4333"))
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	subject := "orders"
	queue := "orders"

	subscribe, err := sc.QueueSubscribe(subject, queue, OrdersHandler)

	if err != nil {
		log.Fatal(err)
	}

	defer subscribe.Unsubscribe()

	fmt.Printf("Subscribed: %s\n", subject)

	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-channel
}

func OrdersHandler(message *stan.Msg) {
	fmt.Printf("Received: %s\n", string(message.Data))
}
