package orders

import (
	"encoding/json"
	"fmt"
	"github.com/emptyhopes/orders-publisher/internal/repository"
	def "github.com/emptyhopes/orders-publisher/internal/service"
	"github.com/nats-io/stan.go"
	"log"
)

type service struct {
	orderRepository repository.OrdersRepositoryInterface
}

var _ def.OrdersServiceInterface = &service{}

func NewService(orderRepository repository.OrdersRepositoryInterface) *service {
	return &service{
		orderRepository: orderRepository,
	}
}

func (s *service) PublishOrder(sc stan.Conn, subject string) {
	order := s.orderRepository.GetOrder()

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
