package controller

import "github.com/nats-io/stan.go"

type OrderControllerInterface interface {
	PublishOrder(sc stan.Conn, subject string)
}
