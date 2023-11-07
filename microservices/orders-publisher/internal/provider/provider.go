package provider

import (
	"github.com/emptyhopes/orders-publisher/internal/controller"
	"github.com/emptyhopes/orders-publisher/internal/repository"
	"github.com/emptyhopes/orders-publisher/internal/service"
)

type OrderProviderInterface interface {
	GetOrderController() controller.OrderControllerInterface
	GetOrderService() service.OrderServiceInterface
	GetOrderRepository() repository.OrderRepositoryInterface
}
