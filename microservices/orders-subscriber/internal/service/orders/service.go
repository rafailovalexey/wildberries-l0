package orders

import (
	"fmt"
	dto "github.com/emptyhopes/orders-subscriber/internal/dto/orders"
	repository "github.com/emptyhopes/orders-subscriber/internal/repository/orders"
	def "github.com/emptyhopes/orders-subscriber/internal/service"
)

type service struct{}

var _ def.OrdersServiceInterface = &service{}

func NewService() *service {
	return &service{}
}

func (s *service) HandleOrderMessage(order *dto.OrderDto) {
	repositoryOrders := repository.NewRepository()

	repositoryOrders.SetOrderCache(order.OrderUid, order)

	ordersCache := repositoryOrders.GetOrdersCache()

	for _, value := range ordersCache {
		orderDto, ok := value.Data.(*dto.OrderDto)

		if !ok {
			fmt.Printf("ошибка при приведение типа")

			repositoryOrders.DeleteOrderCacheById(orderDto.OrderUid)

			return
		}

		err := repositoryOrders.CreateOrder(orderDto)

		if err != nil {
			fmt.Printf("ошибка при создание заказа %v\n", err)

			return
		}

		fmt.Printf("обработал сообщение с order_uid: %s\n", order.OrderUid)

		repositoryOrders.DeleteOrderCacheById(orderDto.OrderUid)

		fmt.Printf("очистил кэш для order_uid: %s\n", order.OrderUid)
	}
}
