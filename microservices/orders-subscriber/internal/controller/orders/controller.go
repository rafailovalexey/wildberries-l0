package orders

import (
	"encoding/json"
	"fmt"
	definition "github.com/emptyhopes/orders-subscriber/internal/controller"
	dto "github.com/emptyhopes/orders-subscriber/internal/dto/orders"
	"github.com/emptyhopes/orders-subscriber/internal/service"
	"github.com/nats-io/stan.go"
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

func (c *controller) HandleOrderMessage(message *stan.Msg) {
	var order dto.OrderDto

	err := json.Unmarshal(message.Data, &order)

	if err != nil {
		fmt.Printf("произошла ошибка парсинга %v\n", err)

		return
	}

	c.orderService.HandleOrderMessage(&order)
}
