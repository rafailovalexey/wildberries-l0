package service

import (
	model "github.com/emptyhopes/level0/internal/model/orders"
)

type OrdersServiceInterface interface {
	GetOrderById(id string) (*model.OrderModel, error)
}
