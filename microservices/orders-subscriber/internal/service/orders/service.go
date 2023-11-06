package orders

import (
	"encoding/json"
	"fmt"
	dto "github.com/emptyhopes/orders-subscriber/internal/dto/orders"
	repository "github.com/emptyhopes/orders-subscriber/internal/repository/orders"
	"github.com/emptyhopes/orders-subscriber/internal/service"
	"github.com/nats-io/stan.go"
	"log"
)

type Service struct{}

var _ service.OrdersServiceInterface = &Service{}

func (s *Service) SubscribeOrders(message *stan.Msg) {
	var data dto.OrderDto

	err := json.Unmarshal(message.Data, &data)

	if err != nil {
		log.Fatalf("ошибка %v\n", err)
	}

	repositoryOrders := &repository.Repository{}

	repositoryOrders.SetOrderCache(data.OrderUid, &data)

	ordersCache := repositoryOrders.GetOrdersCache()

	for _, value := range ordersCache {
		orderDto, ok := value.Data.(*dto.OrderDto)

		if !ok {
			log.Fatalf("ошибка при приведение типа")
		}

		err = repositoryOrders.CreateOrder(orderDto)

		if err != nil {
			log.Fatalf("ошибка при создание заказа %v\n", err)
		}

		fmt.Printf("обработал сообщение с order_uid: %s\n", data.OrderUid)

		repositoryOrders.DeleteOrderCacheById(orderDto.OrderUid)

		fmt.Printf("очистил кэш для order_uid: %s\n", data.OrderUid)
	}
}
