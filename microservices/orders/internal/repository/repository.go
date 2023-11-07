package repository

import (
	dto "github.com/emptyhopes/orders/internal/dto/orders"
)

type OrderRepositoryInterface interface {
	GetOrderCacheById(string) (*dto.OrderDto, bool)
	SetOrderCache(string, *dto.OrderDto)
	GetOrderById(string) (*dto.OrderDto, error)
}
