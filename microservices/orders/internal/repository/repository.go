package repository

import (
	dto "github.com/emptyhopes/orders/internal/dto/orders"
	"github.com/emptyhopes/orders/storage"
)

var Cache = storage.ConstructorCache()
var Database = storage.ConstructorDatabase()

func init() {
	Database.Initialize()
}

type OrdersRepositoryInterface interface {
	GetOrderCache(string) (*dto.OrderDto, bool)
	SetOrderCache(string, *dto.OrderDto)
	GetOrderById(string) (*dto.OrderDto, error)
}
