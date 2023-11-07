package service

import (
	"github.com/nats-io/stan.go"
)

type OrdersServiceInterface interface {
	PublishOrder(sc stan.Conn, subject string)
}
