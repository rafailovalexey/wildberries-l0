package service

import (
	model "github.com/emptyhopes/orders/internal/model/orders"
)

type OrdersServiceInterface interface {
	GetOrderById(id string) (*model.OrderModel, error)
}
