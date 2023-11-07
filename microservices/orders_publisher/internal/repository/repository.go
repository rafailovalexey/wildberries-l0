package repository

import (
	dto "github.com/emptyhopes/orders_publisher/internal/dto/orders"
)

type OrderRepositoryInterface interface {
	GetOrder() *dto.OrderDto
}
