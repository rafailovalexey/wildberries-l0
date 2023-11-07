package repository

import (
	dto "github.com/emptyhopes/orders-publisher/internal/dto/orders"
)

type OrderRepositoryInterface interface {
	GetOrder() *dto.OrderDto
}
