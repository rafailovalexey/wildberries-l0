package nats_subscriber

import (
	"fmt"
	service "github.com/emptyhopes/orders-subscriber/internal/service/orders"
	"github.com/nats-io/stan.go"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Start() {
	url := os.Getenv("NATS_URL")

	if url == "" {
		log.Fatalf("укажите nats-publisher url")
	}

	cluster := os.Getenv("NATS_CLUSTER_ID")

	if cluster == "" {
		log.Fatalf("укажите идентификатор кластера")
	}

	sc, err := stan.Connect(cluster, "subscriber-1", stan.NatsURL(url))

	if err != nil {
		log.Fatal(err)
	}

	defer sc.Close()

	serviceOrders := &service.Service{}

	Subscribe(sc, "orders", "orders", serviceOrders.SubscribeOrders)
}

func Subscribe(sc stan.Conn, subject string, queue string, handler stan.MsgHandler) {
	subscribe, err := sc.QueueSubscribe(subject, queue, handler)

	if err != nil {
		log.Fatal(err)
	}

	defer subscribe.Unsubscribe()

	fmt.Printf("Subscribed: %s\n", subject)

	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-channel
}
