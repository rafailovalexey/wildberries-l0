package repository

import (
	dto "github.com/emptyhopes/orders_subscriber/internal/dto/orders"
	"github.com/emptyhopes/orders_subscriber/storage"
)

type OrderRepositoryInterface interface {
	GetOrdersCache() *map[string]storage.CacheItem
	GetOrderCacheById(string) (*dto.OrderDto, bool)
	SetOrderCache(string, *dto.OrderDto)
	DeleteOrderCacheById(id string)
	CreateOrder(*dto.OrderDto) error
}
