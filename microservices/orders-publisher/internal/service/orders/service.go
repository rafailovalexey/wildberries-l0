package orders

import (
	"fmt"
	"github.com/emptyhopes/orders-publisher/internal/service"
	"github.com/nats-io/stan.go"
	"log"
	"time"
)

type Service struct{}

var _ service.OrdersServiceInterface = &Service{}

func (s *Service) PublishOrders(sc stan.Conn, subject string) {
	message := "ya tyt"

	for {
		err := sc.Publish(subject, []byte(message))

		if err != nil {
			log.Printf("Error: %v\n", err)
		}

		if err == nil {
			fmt.Printf("Published: %s\n", message)
		}

		time.Sleep(1000 * time.Millisecond)
	}
}
