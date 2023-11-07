package orders

import (
	"fmt"
	dto "github.com/emptyhopes/orders_subscriber/internal/dto/orders"
	"github.com/emptyhopes/orders_subscriber/internal/repository"
	definition "github.com/emptyhopes/orders_subscriber/internal/service"
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
	fmt.Printf("добавил в кэш сообщение с order_uid: %s\n", order.OrderUid)

	s.orderRepository.SetOrderCache(order.OrderUid, order)

	ordersCache := s.orderRepository.GetOrdersCache()

	for _, value := range ordersCache {
		orderDto, isExist := value.Data.(*dto.OrderDto)

		if !isExist {
			fmt.Printf("ошибка при приведение типа")

			s.orderRepository.DeleteOrderCacheById(orderDto.OrderUid)

			return
		}

		err := s.orderRepository.CreateOrder(orderDto)

		if err != nil {
			fmt.Printf("ошибка при создание заказа %v\n", err)

			return
		}

		fmt.Printf("обработал сообщение с order_uid: %s\n", order.OrderUid)

		s.orderRepository.DeleteOrderCacheById(orderDto.OrderUid)

		fmt.Printf("удалил из кэша сообщение с order_uid: %s\n", order.OrderUid)
	}
}
