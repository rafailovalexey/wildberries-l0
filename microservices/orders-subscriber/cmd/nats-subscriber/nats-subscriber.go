package nats_subscriber

import (
	"fmt"
	"github.com/emptyhopes/orders-subscriber/internal/controller"
	"github.com/nats-io/stan.go"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Start(orderController controller.OrderControllerInterface) {
	sc := Connect()

	defer sc.Close()

	Subscribe(sc, "orders", "orders", orderController.HandleOrderMessage)
}

func Connect() stan.Conn {
	url := os.Getenv("NATS_URL")

	if url == "" {
		log.Fatalf("укажите nats-publisher url")
	}

	cluster := os.Getenv("NATS_CLUSTER_ID")

	if cluster == "" {
		log.Fatalf("укажите идентификатор кластера")
	}

	sc, err := stan.Connect(cluster, "publisher-1", stan.NatsURL(url))

	if err != nil {
		log.Fatalf("ошибка %v\n", err)
	}

	return sc
}

func Subscribe(sc stan.Conn, subject string, queue string, handler stan.MsgHandler) {
	subscribe, err := sc.QueueSubscribe(subject, queue, handler)

	if err != nil {
		log.Fatalf("ошибка %v\n", err)
	}

	defer subscribe.Unsubscribe()

	fmt.Printf("подписался на очередь сообщений: %s\n", subject)

	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-channel
}
