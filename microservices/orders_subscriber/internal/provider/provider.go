package provider

import (
	"github.com/emptyhopes/orders_subscriber/internal/controller"
	"github.com/emptyhopes/orders_subscriber/internal/converter"
	"github.com/emptyhopes/orders_subscriber/internal/repository"
	"github.com/emptyhopes/orders_subscriber/internal/service"
)

type OrderProviderInterface interface {
	GetOrderController() controller.OrderControllerInterface
	GetOrderService() service.OrderServiceInterface
	GetOrderRepository() repository.OrderRepositoryInterface
	GetOrderConverter() converter.OrderConverterInterface
}
