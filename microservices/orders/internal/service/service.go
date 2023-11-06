package service

import (
	dto "github.com/emptyhopes/orders/internal/dto/orders"
)

type OrdersServiceInterface interface {
	GetOrderById(id string) (*dto.OrderDto, error)
}
