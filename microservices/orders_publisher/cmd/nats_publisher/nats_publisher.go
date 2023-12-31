package nats_publisher

import (
	"github.com/emptyhopes/orders_publisher/internal/controller"
	"github.com/nats-io/stan.go"
	"log"
	"os"
	"time"
)

func Start(orderController controller.OrderControllerInterface) {
	sc := connect()

	defer sc.Close()

	for {
		orderController.PublishOrder(sc, "orders")

		time.Sleep(10 * time.Second)
	}
}

func connect() stan.Conn {
	url := os.Getenv("NATS_URL")

	if url == "" {
		log.Panicf("specify nats url")
	}

	cluster := os.Getenv("NATS_CLUSTER_ID")

	if cluster == "" {
		log.Panicf("specify the cluster id")
	}

	sc, err := stan.Connect(cluster, "publisher-1", stan.NatsURL(url))

	if err != nil {
		log.Panicf("error %v\n", err)
	}

	return sc
}
