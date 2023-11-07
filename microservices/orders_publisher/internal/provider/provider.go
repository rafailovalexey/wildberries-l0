package provider

import (
	"github.com/emptyhopes/orders_publisher/internal/controller"
	"github.com/emptyhopes/orders_publisher/internal/repository"
	"github.com/emptyhopes/orders_publisher/internal/service"
)

type OrderProviderInterface interface {
	GetOrderController() controller.OrderControllerInterface
	GetOrderService() service.OrderServiceInterface
	GetOrderRepository() repository.OrderRepositoryInterface
}
