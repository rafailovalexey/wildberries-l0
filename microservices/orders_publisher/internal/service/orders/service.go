package orders

import (
	"encoding/json"
	"github.com/emptyhopes/orders_publisher/internal/repository"
	definition "github.com/emptyhopes/orders_publisher/internal/service"
	"github.com/nats-io/stan.go"
	"log"
)

type service struct {
	orderRepository repository.OrderRepositoryInterface
}

var _ definition.OrderServiceInterface = (*service)(nil)

func NewOrderService(orderRepository repository.OrderRepositoryInterface) *service {
	return &service{
		orderRepository: orderRepository,
	}
}

func (s *service) PublishOrder(sc stan.Conn, subject string) {
	order := s.orderRepository.GetOrder()

	message, err := json.Marshal(order)

	if err != nil {
		log.Panicf("error %v\n", err)
	}

	err = sc.Publish(subject, []byte(message))

	if err != nil {
		log.Panicf("error %v\n", err)
	}

	if err == nil {
		log.Printf("posted a message with order_uid %s\n", order.OrderUid)
	}
}
