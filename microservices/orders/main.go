package main

import (
	"context"
	"github.com/emptyhopes/orders/cmd/application"
	"log"
)

func main() {
	ctx := context.Background()

	app, err := application.NewApplication(ctx)

	if err != nil {
		log.Panicf("an error occurred during initialization %v", err)
	}

	app.Run()
}
