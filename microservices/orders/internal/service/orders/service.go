package orders

import (
	"fmt"
	dto "github.com/emptyhopes/orders/internal/dto/orders"
	repository "github.com/emptyhopes/orders/internal/repository/orders"
	"github.com/emptyhopes/orders/internal/service"
)

type Service struct{}

var _ service.OrdersServiceInterface = &Service{}

func (s *Service) GetOrderById(id string) (*dto.OrderDto, error) {
	orderRepository := &repository.Repository{}

	orderDto, err := orderRepository.GetOrderById(id)
	if err != nil {
		fmt.Println("ya tyt", err)
		return nil, err
	}

	return orderDto, nil
}
