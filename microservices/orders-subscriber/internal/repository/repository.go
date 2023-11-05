package repository

import (
	dto "github.com/emptyhopes/orders-subscriber/internal/dto/orders"
	"github.com/emptyhopes/orders-subscriber/storage"
)

var Cache = storage.ConstructorCache()
var Database = storage.ConstructorDatabase()

func init() {
	Database.Initialize()

	pool := Database.GetPool()
	defer pool.Close()

	Database.CreateTables(pool)
}

type OrdersRepositoryInterface interface {
	Cache(*dto.OrderDto) *dto.OrderDto
	CreateOrder(*dto.OrderDto) error
}
