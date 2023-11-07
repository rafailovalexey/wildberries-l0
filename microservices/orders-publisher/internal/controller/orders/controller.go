package orders

import (
	definition "github.com/emptyhopes/orders-publisher/internal/controller"
	"github.com/emptyhopes/orders-publisher/internal/service"
	"github.com/nats-io/stan.go"
	"time"
)

type controller struct {
	orderService service.OrderServiceInterface
}

var _ definition.OrderControllerInterface = (*controller)(nil)

func NewOrderController(orderService service.OrderServiceInterface) *controller {
	return &controller{
		orderService: orderService,
	}
}

func (c *controller) PublishOrder(sc stan.Conn, subject string) {
	for {
		c.orderService.PublishOrder(sc, subject)

		time.Sleep(10 * time.Second)
	}
}
