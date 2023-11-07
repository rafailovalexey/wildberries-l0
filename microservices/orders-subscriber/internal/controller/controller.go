package controller

import "github.com/nats-io/stan.go"

type OrdersControllerInterface interface {
	HandleOrderMessage(*stan.Msg)
}
