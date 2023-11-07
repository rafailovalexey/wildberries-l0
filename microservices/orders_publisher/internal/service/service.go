package service

import (
	"github.com/nats-io/stan.go"
)

type OrderServiceInterface interface {
	PublishOrder(sc stan.Conn, subject string)
}
