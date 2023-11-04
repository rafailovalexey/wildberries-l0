package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/stan.go"
)

func main() {
	sc, err := stan.Connect("test-cluster", "publisher-1", stan.NatsURL("http://localhost:4333"))
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	subject := "orders"
	message := "ya tyt"

	for {
		err = sc.Publish(subject, []byte(message))
		if err != nil {
			log.Printf("Error: %v\n", err)
		}

		if err == nil {
			fmt.Printf("Published: %s\n", message)
		}

		time.Sleep(1000 * time.Millisecond)
	}
}
