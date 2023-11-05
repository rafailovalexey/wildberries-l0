package orders

import (
	"encoding/json"
	"fmt"
	converter "github.com/emptyhopes/orders-subscriber/internal/converter/orders"
	dto "github.com/emptyhopes/orders-subscriber/internal/dto/orders"
	repository "github.com/emptyhopes/orders-subscriber/internal/repository/orders"
	"github.com/emptyhopes/orders-subscriber/internal/service"
	"github.com/nats-io/stan.go"
	"log"
)

type Service struct {
}

var _ service.OrdersServiceInterface = &Service{}

func (s *Service) SubscribeOrders(message *stan.Msg) {
	var data dto.OrderDto

	err := json.Unmarshal(message.Data, &data)

	if err != nil {
		log.Fatalf("error %v\n", err)
	}

	converterOrders := &converter.Converter{}

	model := converterOrders.OrderDtoToOrderModel(&data)

	repositoryOrders := &repository.Repository{}

	//repositoryOrders.Cache(model)
	repositoryOrders.CreateOrder(model)

	fmt.Printf("%#v\n", model)
}
