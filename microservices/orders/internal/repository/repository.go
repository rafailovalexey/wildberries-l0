package repository

import (
	dto "github.com/emptyhopes/orders/internal/dto/orders"
	"github.com/emptyhopes/orders/storage"
)

var Cache = storage.NewCache()
var Database = storage.NewDatabase()

func init() {
	Database.Initialize()
}

type OrdersRepositoryInterface interface {
	GetOrderCache(string) (*dto.OrderDto, bool)
	SetOrderCache(string, *dto.OrderDto)
	GetOrderById(string) (*dto.OrderDto, error)
}
