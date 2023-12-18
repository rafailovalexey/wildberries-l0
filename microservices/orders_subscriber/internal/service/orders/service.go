package orders

import (
	dto "github.com/emptyhopes/orders_subscriber/internal/dto/orders"
	"github.com/emptyhopes/orders_subscriber/internal/repository"
	definition "github.com/emptyhopes/orders_subscriber/internal/service"
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

func (s *service) HandleOrderMessage(order *dto.OrderDto) {
	log.Printf("added a message with order_uid to the cache %s\n", order.OrderUid)

	s.orderRepository.SetOrderCache(order.OrderUid, order)

	ordersCache := s.orderRepository.GetOrdersCache()

	for _, value := range *ordersCache {
		orderDto, isExist := value.Data.(*dto.OrderDto)

		if !isExist {
			log.Printf("error when casting type")

			s.orderRepository.DeleteOrderCacheById(orderDto.OrderUid)

			return
		}

		err := s.orderRepository.CreateOrder(orderDto)

		if err != nil {
			log.Printf("error when creating an order %v\n", err)

			return
		}

		log.Printf("processed a message with order_uid %s\n", order.OrderUid)

		s.orderRepository.DeleteOrderCacheById(orderDto.OrderUid)

		log.Printf("deleted the message with order_uid from the cache %s\n", order.OrderUid)
	}
}
