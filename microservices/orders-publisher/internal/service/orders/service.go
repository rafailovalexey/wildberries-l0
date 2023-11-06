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
		uuid.New().String(),
		"1",
		"USD",
		"wbpay",
		1817,
		time.Now().Unix(),
		"alpha",
		1500,
		317,
		0,
	)
	item1 := dto.ConstructorOrderItemDto(
		"WBILMTESTTRACK",
		453,
		uuid.New().String(),
		"Mascaras",
		30,
		"0",
		317,
		2389212,
		"Vivienne Sabo",
		202,
	)
	item2 := dto.ConstructorOrderItemDto(
		"WBILMTESTTRACK",
		453,
		uuid.New().String(),
		"Mascaras",
		30,
		"0",
		317,
		2389212,
		"Vivienne Sabo",
		202,
	)
	item3 := dto.ConstructorOrderItemDto(
		"WBILMTESTTRACK",
		453,
		uuid.New().String(),
		"Mascaras",
		30,
		"0",
		317,
		2389212,
		"Vivienne Sabo",
		202,
	)
	item4 := dto.ConstructorOrderItemDto(
		"WBILMTESTTRACK",
		453,
		uuid.New().String(),
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
		*item2,
		*item3,
		*item4,
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
			time.Now().Unix(),
			"1",
		)

		message, err := json.Marshal(order)

		if err != nil {
			log.Fatalf("ошибка %v\n", err)
		}

		err = sc.Publish(subject, []byte(message))

		if err != nil {
			log.Fatalf("ошибка %v\n", err)
		}

		if err == nil {
			fmt.Printf("опубликовал сообщение с order_uid: %s\n", order.OrderUid)
		}

		time.Sleep(10 * time.Second)
	}
}
