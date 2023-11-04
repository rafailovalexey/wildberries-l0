package orders

import (
	"encoding/json"
	"fmt"
	"github.com/emptyhopes/orders-publisher/internal/model/orders"
	"github.com/emptyhopes/orders-publisher/internal/service"
	"github.com/google/uuid"
	"github.com/nats-io/stan.go"
	"log"
	"time"
)

type Service struct{}

var _ service.OrdersServiceInterface = &Service{}

func (s *Service) PublishOrders(sc stan.Conn, subject string) {
	delivery := orders.ConstructorOrderDeliveryModel(
		"",
		"",
		"",
		"",
		"",
		"",
		"",
	)
	payment := orders.ConstructorOrderPaymentModel(
		"",
		"",
		"",
		"",
		0,
		0,
		"",
		0,
		0,
		0,
	)
	item1 := orders.ConstructorOrderItemModel(
		0,
		"",
		0,
		"",
		"",
		0,
		"",
		0,
		0,
		"",
		0,
	)
	item2 := orders.ConstructorOrderItemModel(
		0,
		"",
		0,
		"",
		"",
		0,
		"",
		0,
		0,
		"",
		0,
	)
	items := &[]orders.OrderItemModel{
		*item1,
		*item2,
	}

	for {
		order := orders.ConstructorOrderModel(
			uuid.New().String(),
			"",
			"",
			*delivery,
			*payment,
			*items,
			"",
			"",
			"",
			"",
			"",
			0,
			time.Now(),
			"",
		)

		message, err := json.Marshal(order)

		if err != nil {
			log.Printf("Error: %v\n", err)
		}

		err = sc.Publish(subject, []byte(message))

		if err != nil {
			log.Printf("Error: %v\n", err)
		}

		if err == nil {
			fmt.Printf("Published: %s\n", message)
		}

		time.Sleep(1000 * time.Millisecond)
	}
}
