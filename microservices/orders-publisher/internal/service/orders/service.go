package orders

import (
	"encoding/json"
	"fmt"
	repository "github.com/emptyhopes/orders-publisher/internal/repository/orders"
	def "github.com/emptyhopes/orders-publisher/internal/service"
	"github.com/nats-io/stan.go"
	"log"
)

type service struct{}

var _ def.OrdersServiceInterface = &service{}

func NewService() *service {
	return &service{}
}

func (s *service) PublishOrder(sc stan.Conn, subject string) {
	orderRepository := repository.NewRepository()

	order := orderRepository.GetOrder()

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
}
