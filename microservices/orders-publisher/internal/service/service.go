package service

import "github.com/nats-io/stan.go"

type OrdersServiceInterface interface {
	PublishOrders(sc stan.Conn, subject string)
}
