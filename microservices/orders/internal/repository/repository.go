package repository

import (
	dto "github.com/emptyhopes/orders/internal/dto/orders"
)

//var Cache = storage.NewCache()
//var Database = storage.NewDatabase()

//func init() {
//	Database.Initialize()
//}

type OrderRepositoryInterface interface {
	GetOrderCache(string) (*dto.OrderDto, bool)
	SetOrderCache(string, *dto.OrderDto)
	GetOrderById(string) (*dto.OrderDto, error)
}
