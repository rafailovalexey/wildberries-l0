package provider

import (
	"github.com/emptyhopes/orders/internal/api"
	"github.com/emptyhopes/orders/internal/converter"
	"github.com/emptyhopes/orders/internal/repository"
	"github.com/emptyhopes/orders/internal/service"
)

type OrderProviderInterface interface {
	GetOrderApi() api.OrderApiInterface
	GetOrderService() service.OrderServiceInterface
	GetOrderRepository() repository.OrderRepositoryInterface
	GetOrderConverter() converter.OrderConverterInterface
}
