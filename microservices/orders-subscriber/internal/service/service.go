package service

import (
	"github.com/nats-io/stan.go"
)

type OrdersServiceInterface interface {
	SubscribeOrders(message *stan.Msg)
}
