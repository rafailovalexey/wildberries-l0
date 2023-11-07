package orders

import (
	"github.com/emptyhopes/orders/internal/api"
	orderApi "github.com/emptyhopes/orders/internal/api/orders"
	"github.com/emptyhopes/orders/internal/converter"
	orderConverter "github.com/emptyhopes/orders/internal/converter/orders"
	defenition "github.com/emptyhopes/orders/internal/provider"
	"github.com/emptyhopes/orders/internal/repository"
	orderRepository "github.com/emptyhopes/orders/internal/repository/orders"
	"github.com/emptyhopes/orders/internal/service"
	orderService "github.com/emptyhopes/orders/internal/service/orders"
	"github.com/emptyhopes/orders/storage"
)

type provider struct {
	orderApi        api.OrderApiInterface
	orderService    service.OrderServiceInterface
	orderRepository repository.OrderRepositoryInterface
	orderConverter  converter.OrderConverterInterface
}

var _ defenition.OrderProviderInterface = (*provider)(nil)

func NewOrderProvider() *provider {
	return &provider{}
}

func (p *provider) GetOrderApi() api.OrderApiInterface {
	if p.orderApi == nil {
		p.orderApi = orderApi.NewOrderApi(
			p.GetOrderService(),
		)
	}

	return p.orderApi
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
