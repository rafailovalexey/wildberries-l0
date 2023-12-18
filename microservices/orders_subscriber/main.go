package main

import (
	"context"
	"github.com/emptyhopes/orders_subscriber/cmd/subscriber"
	"log"
)

func main() {
	ctx := context.Background()

	sub, err := subscriber.NewSubscriber(ctx)

	if err != nil {
		log.Panicf("произошла ошибка при инициализации %v", err)
	}

	sub.Run()
}
