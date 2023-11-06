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
	GetOrderCache(string) (*dto.OrderDto, bool)
	SetOrderCache(string, *dto.OrderDto)
	CreateOrder(*dto.OrderDto) error
}
