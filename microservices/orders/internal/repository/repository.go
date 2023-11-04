package repository

import (
	model "github.com/emptyhopes/orders/internal/model/orders"
	"github.com/emptyhopes/orders/storage"
)

var Cache = storage.ConstructorCache()

type OrdersRepositoryInterface interface {
	GetOrderById(id string) (*model.OrderModel, error)
}
