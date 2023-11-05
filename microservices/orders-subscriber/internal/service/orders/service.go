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
		log.Fatalf("error %v\n", err)
	}

	repositoryOrders := &repository.Repository{}

	repositoryOrders.CreateOrder(&data)

	fmt.Println()
}
