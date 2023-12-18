package orders

import (
	"encoding/json"
	definition "github.com/emptyhopes/orders_subscriber/internal/controller"
	dto "github.com/emptyhopes/orders_subscriber/internal/dto/orders"
	"github.com/emptyhopes/orders_subscriber/internal/service"
	"github.com/emptyhopes/orders_subscriber/internal/validation"
	"github.com/nats-io/stan.go"
	"log"
)

type controller struct {
	orderValidation validation.OrderValidationInterface
	orderService    service.OrderServiceInterface
}

var _ definition.OrderControllerInterface = (*controller)(nil)

func NewOrderController(orderValidation validation.OrderValidationInterface, orderService service.OrderServiceInterface) *controller {
	return &controller{
		orderValidation: orderValidation,
		orderService:    orderService,
	}
}

func (c *controller) HandleOrderMessage(message *stan.Msg) {
	var order dto.OrderDto

	err := json.Unmarshal(message.Data, &order)

	if err != nil {
		log.Printf("a parsing error occurred %v\n", err)

		return
	}

	err = c.orderValidation.HandleOrderMessageValidation(&order)

	if err != nil {
		log.Printf("a validation error occurred %v\n", err)

		return
	}

	c.orderService.HandleOrderMessage(&order)
}
