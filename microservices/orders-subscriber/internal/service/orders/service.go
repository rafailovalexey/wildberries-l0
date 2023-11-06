package orders

import (
	"encoding/json"
	"fmt"
	dto "github.com/emptyhopes/orders-subscriber/internal/dto/orders"
	repository "github.com/emptyhopes/orders-subscriber/internal/repository/orders"
	"github.com/emptyhopes/orders-subscriber/internal/service"
	"github.com/nats-io/stan.go"
	"log"
)

type Service struct{}

var _ service.OrdersServiceInterface = &Service{}

func (s *Service) SubscribeOrders(message *stan.Msg) {
	var data dto.OrderDto

	err := json.Unmarshal(message.Data, &data)

	if err != nil {
		log.Fatalf("ошибка %v\n", err)
	}

	repositoryOrders := &repository.Repository{}

	err = repositoryOrders.CreateOrder(&data)

	if err != nil {
		log.Fatalf("ошибка %v\n", err)
	}

	fmt.Printf("обработал сообщение с order_uid: %s\n", data.OrderUid)
}
