package nats_subscriber

import (
	"fmt"
	"github.com/emptyhopes/orders_subscriber/internal/controller"
	"github.com/nats-io/stan.go"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Start(orderController controller.OrderControllerInterface) {
	sc := connect()

	defer sc.Close()

	subscribe(sc, "orders", "orders", orderController.HandleOrderMessage)
}

func connect() stan.Conn {
	url := os.Getenv("NATS_URL")

	if url == "" {
		log.Fatalf("укажите nats url")
	}

	cluster := os.Getenv("NATS_CLUSTER_ID")

	if cluster == "" {
		log.Fatalf("укажите идентификатор кластера")
	}

	sc, err := stan.Connect(cluster, "subscriber-1", stan.NatsURL(url))

	if err != nil {
		log.Fatalf("ошибка %v\n", err)
	}

	return sc
}

func subscribe(sc stan.Conn, subject string, queue string, handler stan.MsgHandler) {
	sub, err := sc.QueueSubscribe(subject, queue, handler)

	if err != nil {
		log.Fatalf("ошибка %v\n", err)
	}

	defer sub.Unsubscribe()

	fmt.Printf("подписался на очередь сообщений: %s\n", subject)

	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-channel
}
