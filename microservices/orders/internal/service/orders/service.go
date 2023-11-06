package orders

import (
	"errors"
	"fmt"
	dto "github.com/emptyhopes/orders/internal/dto/orders"
	repository "github.com/emptyhopes/orders/internal/repository/orders"
	"github.com/emptyhopes/orders/internal/service"
	"strings"
)

type Service struct{}

var _ service.OrdersServiceInterface = &Service{}

func (s *Service) GetOrderById(id string) (*dto.OrderDto, error) {
	orderRepository := &repository.Repository{}

	orderCached, isExist := orderRepository.GetOrderCache(id)

	if isExist {
		fmt.Printf("пользователь получил данные из кэша по заказу с order_uid: %s\n", orderCached.OrderUid)

		return orderCached, nil
	}

	orderDto, err := orderRepository.GetOrderById(id)

	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return nil, errors.New(fmt.Sprintf("пользователь не найден с order_uid: %s\n", orderDto.OrderUid))
		}

		return nil, err
	}

	orderRepository.SetOrderCache(orderDto.OrderUid, orderDto)

	fmt.Printf("пользователь получил данные из базы данных по заказу с order_uid: %s\n", orderDto.OrderUid)

	return orderDto, nil
}
