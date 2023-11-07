package main

import (
	"context"
	"github.com/emptyhopes/orders-publisher/cmd/publisher"
	"log"
)

func main() {
	ctx := context.Background()

	pub, err := publisher.NewApplication(ctx)

	if err != nil {
		log.Fatalf("произошла ошибка при инициализации %v", err)
	}

	pub.Run()
}
