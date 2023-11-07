package service

import (
	dto "github.com/emptyhopes/orders_subscriber/internal/dto/orders"
)

type OrderServiceInterface interface {
	HandleOrderMessage(*dto.OrderDto)
}
