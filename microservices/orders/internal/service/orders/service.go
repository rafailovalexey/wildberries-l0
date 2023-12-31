package orders

import (
	dto "github.com/emptyhopes/orders/internal/dto/orders"
	"github.com/emptyhopes/orders/internal/repository"
	definition "github.com/emptyhopes/orders/internal/service"
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

func (s *service) GetOrderById(id string) (*dto.OrderDto, error) {
	orderCached, isExist := s.orderRepository.GetOrderCacheById(id)

	if isExist {
		log.Printf("the user received data from the cache by order with order_uid %s\n", orderCached.OrderUid)

		return orderCached, nil
	}

	orderDto, err := s.orderRepository.GetOrderById(id)

	if err != nil {
		return nil, err
	}

	s.orderRepository.SetOrderCache(orderDto.OrderUid, orderDto)

	log.Printf("the user received data from the database for an order with order_uid %s\n", orderDto.OrderUid)

	return orderDto, nil
}
