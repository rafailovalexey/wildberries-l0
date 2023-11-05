package repository

import (
	model "github.com/emptyhopes/orders-subscriber/internal/model/orders"
	"github.com/emptyhopes/orders-subscriber/storage"
)

var Cache = storage.ConstructorCache()

func init() {
	database := &storage.Database{}

	database.Initialize()

	pool := database.GetPool()
	defer pool.Close()

	database.CreateTables(pool)
}

type OrdersRepositoryInterface interface {
	Cache(*model.OrderModel) (bool, error)
	CreateOrder(order *model.OrderModel) error
}
