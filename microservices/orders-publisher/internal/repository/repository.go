package repository

import (
	dto "github.com/emptyhopes/orders-publisher/internal/dto/orders"
)

type OrdersRepositoryInterface interface {
	GetOrder() *dto.OrderDto
}
