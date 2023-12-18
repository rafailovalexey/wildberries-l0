package main

import (
	"context"
	"github.com/emptyhopes/orders_publisher/cmd/publisher"
	"log"
)

func main() {
	ctx := context.Background()

	pub, err := publisher.NewPublisher(ctx)

	if err != nil {
		log.Panicf("произошла ошибка при инициализации %v", err)
	}

	pub.Run()
}
