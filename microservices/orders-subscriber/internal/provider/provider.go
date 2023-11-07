package provider

import (
	"github.com/emptyhopes/orders-subscriber/internal/controller"
	"github.com/emptyhopes/orders-subscriber/internal/converter"
	"github.com/emptyhopes/orders-subscriber/internal/repository"
	"github.com/emptyhopes/orders-subscriber/internal/service"
)

type OrderProviderInterface interface {
	GetOrderController() controller.OrderControllerInterface
	GetOrderService() service.OrderServiceInterface
	GetOrderRepository() repository.OrderRepositoryInterface
	GetOrderConverter() converter.OrderConverterInterface
}
