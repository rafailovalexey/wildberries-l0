package orders

import (
	"github.com/emptyhopes/orders_publisher/internal/controller"
	orderController "github.com/emptyhopes/orders_publisher/internal/controller/orders"
	defenition "github.com/emptyhopes/orders_publisher/internal/provider"
	"github.com/emptyhopes/orders_publisher/internal/repository"
	orderRepository "github.com/emptyhopes/orders_publisher/internal/repository/orders"
	"github.com/emptyhopes/orders_publisher/internal/service"
	orderService "github.com/emptyhopes/orders_publisher/internal/service/orders"
)

type provider struct {
	orderController controller.OrderControllerInterface
	orderService    service.OrderServiceInterface
	orderRepository repository.OrderRepositoryInterface
}

var _ defenition.OrderProviderInterface = (*provider)(nil)

func NewOrderProvider() *provider {
	return &provider{}
}

func (p *provider) GetOrderController() controller.OrderControllerInterface {
	if p.orderController == nil {
		p.orderController = orderController.NewOrderController(
			p.GetOrderService(),
		)
	}

	return p.orderController
}

func (p *provider) GetOrderService() service.OrderServiceInterface {
	if p.orderService == nil {
		p.orderService = orderService.NewOrderService(
			p.GetOrderRepository(),
		)
	}

	return p.orderService
}

func (p *provider) GetOrderRepository() repository.OrderRepositoryInterface {
	if p.orderRepository == nil {
		p.orderRepository = orderRepository.NewOrderRepository()
	}

	return p.orderRepository
}
