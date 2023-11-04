package repository

import (
	model "github.com/emptyhopes/level0/internal/model/orders"
	"github.com/emptyhopes/level0/storage"
)

var Cache = storage.ConstructorCache()

type OrdersRepositoryInterface interface {
	GetOrderById(id string) (*model.OrderModel, error)
}
