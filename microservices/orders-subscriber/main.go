package main

import (
	"context"
	"github.com/emptyhopes/orders-subscriber/cmd/subscriber"
	"log"
)

func main() {
	ctx := context.Background()

	sub, err := subscriber.NewApplication(ctx)

	if err != nil {
		log.Fatalf("произошла ошибка при инициализации %v", err)
	}

	sub.Run()
}
