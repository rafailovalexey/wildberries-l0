package controller

import "github.com/nats-io/stan.go"

type OrderControllerInterface interface {
	HandleOrderMessage(*stan.Msg)
}
