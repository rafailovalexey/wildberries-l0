package orders

import (
	"fmt"
	dto "github.com/emptyhopes/orders/internal/dto/orders"
	repository "github.com/emptyhopes/orders/internal/repository/orders"
	def "github.com/emptyhopes/orders/internal/service"
)

type service struct{}

var _ def.OrdersServiceInterface = &service{}

func (s *service) GetOrderById(id string) (*dto.OrderDto, error) {
	orderRepository := repository.NewRepository()

	orderCached, isExist := orderRepository.GetOrderCache(id)

	if isExist {
		fmt.Printf("пользователь получил данные из кэша по заказу с order_uid: %s\n", orderCached.OrderUid)

		return orderCached, nil
	}

	orderDto, err := orderRepository.GetOrderById(id)

	if err != nil {
		return nil, err
	}

	orderRepository.SetOrderCache(orderDto.OrderUid, orderDto)

	fmt.Printf("пользователь получил данные из базы данных по заказу с order_uid: %s\n", orderDto.OrderUid)

	return orderDto, nil
}
