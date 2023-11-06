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
	//Cache(id string) (*model.OrderModel, error)
	GetOrderById(id string) (*dto.OrderDto, error)
}
