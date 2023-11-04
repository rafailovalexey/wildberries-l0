package orders

import (
	"encoding/json"
	"fmt"
	"github.com/emptyhopes/orders-subscriber/internal/model/orders"
	"github.com/emptyhopes/orders-subscriber/internal/service"
	"github.com/nats-io/stan.go"
)

type Service struct{}

var _ service.OrdersServiceInterface = &Service{}

func (s *Service) SubscribeOrders(message *stan.Msg) {
	var data orders.OrderModel

	err := json.Unmarshal(message.Data, &data)

	if err != nil {
		fmt.Printf("error %v\n", err)
	}

	fmt.Printf("%v\n", data)
}
