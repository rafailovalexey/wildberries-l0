package service

import (
	dto "github.com/emptyhopes/orders-subscriber/internal/dto/orders"
)

type OrdersServiceInterface interface {
	HandleOrderMessage(*dto.OrderDto)
}
