package nats_publisher

import (
	"github.com/emptyhopes/orders-publisher/internal/service/orders"
	"github.com/nats-io/stan.go"
	"log"
	"os"
)

func Start() {
	//url := os.Getenv("NATS_URL")
	//
	//if url == "" {
	//	log.Fatalf("укажите nats-publisher url")
	//}

	cluster := os.Getenv("NATS_CLUSTER_ID")

	if cluster == "" {
		log.Fatalf("укажите идентификатор кластера")
	}

	//fmt.Println(url, cluster)

	//sc, err := stan.Connect(cluster, "publisher-1", stan.NatsURL(url))
	sc, err := stan.Connect(cluster, "publisher-1")
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	service := &orders.Service{}

	service.PublishOrders(sc, "orders")
}
