package orders

import (
	"encoding/json"
	"fmt"
	def "github.com/emptyhopes/orders-subscriber/internal/controller"
	dto "github.com/emptyhopes/orders-subscriber/internal/dto/orders"
	"github.com/emptyhopes/orders-subscriber/internal/service/orders"
	"github.com/nats-io/stan.go"
)

type controller struct{}

var _ def.OrdersControllerInterface = &controller{}

func NewController() *controller {
	return &controller{}
}

func (c *controller) HandleOrderMessage(message *stan.Msg) {
	var order dto.OrderDto

	err := json.Unmarshal(message.Data, &order)

	if err != nil {
		fmt.Printf("произошла ошибка парсинга %v\n", err)

		return
	}

	orderService := orders.NewService()

	orderService.HandleOrderMessage(&order)
}
