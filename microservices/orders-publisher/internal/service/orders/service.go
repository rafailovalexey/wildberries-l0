package orders

import (
	"encoding/json"
	"fmt"
	dto "github.com/emptyhopes/orders-publisher/internal/dto/orders"
	"github.com/emptyhopes/orders-publisher/internal/service"
	"github.com/google/uuid"
	"github.com/nats-io/stan.go"
	"log"
	"time"
)

type Service struct{}

var _ service.OrdersServiceInterface = &Service{}

func (s *Service) PublishOrders(sc stan.Conn, subject string) {
	delivery := dto.ConstructorOrderDeliveryDto(
		"Test Testov",
		"+9720000000",
		"2639809",
		"Kiryat Mozkin",
		"Ploshad Mira 15",
		"Kraiot",
		"test@gmail.com",
	)
	payment := dto.ConstructorOrderPaymentDto(
		"b563feb7b2b84b6test",
		"1",
		"USD",
		"wbpay",
		1817,
		1637907727,
		"alpha",
		1500,
		317,
		0,
	)
	item1 := dto.ConstructorOrderItemDto(
		9934930,
		"WBILMTESTTRACK",
		453,
		"ab4219087a764ae0btest",
		"Mascaras",
		30,
		"0",
		317,
		2389212,
		"Vivienne Sabo",
		202,
	)
	items := &[]dto.OrderItemDto{
		*item1,
	}

	for {
		order := dto.ConstructorOrderDto(
			uuid.New().String(),
			"WBILMTESTTRACK",
			"WBIL",
			delivery,
			payment,
			items,
			"en",
			"1",
			"test",
			"meest",
			"9",
			99,
			time.Now().String(),
			"1",
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
