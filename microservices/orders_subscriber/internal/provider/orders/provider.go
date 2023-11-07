package orders

import (
	"github.com/emptyhopes/orders_subscriber/internal/controller"
	orderController "github.com/emptyhopes/orders_subscriber/internal/controller/orders"
	"github.com/emptyhopes/orders_subscriber/internal/converter"
	orderConverter "github.com/emptyhopes/orders_subscriber/internal/converter/orders"
	defenition "github.com/emptyhopes/orders_subscriber/internal/provider"
	"github.com/emptyhopes/orders_subscriber/internal/repository"
	orderRepository "github.com/emptyhopes/orders_subscriber/internal/repository/orders"
	"github.com/emptyhopes/orders_subscriber/internal/service"
	orderService "github.com/emptyhopes/orders_subscriber/internal/service/orders"
	"github.com/emptyhopes/orders_subscriber/storage"
)

type provider struct {
	orderController controller.OrderControllerInterface
	orderService    service.OrderServiceInterface
	orderRepository repository.OrderRepositoryInterface
	orderConverter  converter.OrderConverterInterface
}

var _ defenition.OrderProviderInterface = &provider{}

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
		p.orderRepository = orderRepository.NewOrderRepository(
			p.GetOrderConverter(),
			storage.NewDatabase(),
			storage.NewCache(),
		)
	}

	return p.orderRepository
}

func (p *provider) GetOrderConverter() converter.OrderConverterInterface {
	if p.orderConverter == nil {
		p.orderConverter = orderConverter.NewOrderConverter()
	}

	return p.orderConverter
}
