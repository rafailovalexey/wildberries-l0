package orders

import (
	"fmt"
	dto "github.com/emptyhopes/orders/internal/dto/orders"
	"github.com/emptyhopes/orders/internal/repository"
	definition "github.com/emptyhopes/orders/internal/service"
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
		fmt.Printf("пользователь получил данные из кэша по заказу с order_uid: %s\n", orderCached.OrderUid)

		return orderCached, nil
	}

	orderDto, err := s.orderRepository.GetOrderById(id)

	if err != nil {
		return nil, err
	}

	s.orderRepository.SetOrderCache(orderDto.OrderUid, orderDto)

	fmt.Printf("пользователь получил данные из базы данных по заказу с order_uid: %s\n", orderDto.OrderUid)

	return orderDto, nil
}
