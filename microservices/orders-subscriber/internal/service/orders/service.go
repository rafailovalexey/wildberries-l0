package orders

import (
	"encoding/json"
	"fmt"
	dto "github.com/emptyhopes/orders-subscriber/internal/dto/orders"
	repository "github.com/emptyhopes/orders-subscriber/internal/repository/orders"
	"github.com/emptyhopes/orders-subscriber/internal/service"
	"github.com/nats-io/stan.go"
)

type Service struct{}

var _ service.OrdersServiceInterface = &Service{}

func (s *Service) SubscribeOrders(message *stan.Msg) {
	var data dto.OrderDto

	err := json.Unmarshal(message.Data, &data)

	if err != nil {
		fmt.Printf("произошла ошибка парсинга %v\n", err)

		return
	}

	repositoryOrders := &repository.Repository{}

	repositoryOrders.SetOrderCache(data.OrderUid, &data)

	ordersCache := repositoryOrders.GetOrdersCache()

	for _, value := range ordersCache {
		orderDto, ok := value.Data.(*dto.OrderDto)

		if !ok {
			fmt.Printf("ошибка при приведение типа")

			repositoryOrders.DeleteOrderCacheById(orderDto.OrderUid)

			return
		}

		err = repositoryOrders.CreateOrder(orderDto)

		if err != nil {
			fmt.Printf("ошибка при создание заказа %v\n", err)

			return
		}

		fmt.Printf("обработал сообщение с order_uid: %s\n", data.OrderUid)

		repositoryOrders.DeleteOrderCacheById(orderDto.OrderUid)

		fmt.Printf("очистил кэш для order_uid: %s\n", data.OrderUid)
	}
}
