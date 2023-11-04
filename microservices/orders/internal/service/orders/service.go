package orders

import (
	model "github.com/emptyhopes/orders/internal/model/orders"
	repository "github.com/emptyhopes/orders/internal/repository/orders"
	"github.com/emptyhopes/orders/internal/service"
)

type Service struct{}

var _ service.OrdersServiceInterface = &Service{}

func (s *Service) GetOrderById(id string) (*model.OrderModel, error) {
	orderRepository := &repository.Repository{}

	orderRepository.GetOrderById(id)

	return nil, nil
}
