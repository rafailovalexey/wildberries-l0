package orders

import (
	def "github.com/emptyhopes/orders-publisher/internal/controller"
	service "github.com/emptyhopes/orders-publisher/internal/service/orders"
	"github.com/nats-io/stan.go"
	"time"
)

type controller struct{}

var _ def.OrderControllerInterface = &controller{}

func NewController() *controller {
	return &controller{}
}

func (c *controller) PublishOrder(sc stan.Conn, subject string) {
	orderService := service.NewService()

	for {
		orderService.PublishOrder(sc, subject)

		time.Sleep(10 * time.Second)
	}
}
