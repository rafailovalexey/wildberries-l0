package orders

import (
	"encoding/json"
	"fmt"
	def "github.com/emptyhopes/orders/internal/api"
	service "github.com/emptyhopes/orders/internal/service/orders"
	"net/http"
	"strings"
)

type api struct{}

var _ def.OrdersApiInterface = &api{}

/*
OrdersHandler
Использовал парсинг URL, для того, чтобы добиться REST поведения
GetAllOrders - /v1/orders
GetOrderById - /v1/orders/:id
*/
func (a *api) OrdersHandler(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		if strings.HasPrefix(request.URL.Path, "/v1/orders/") {
			a.GetOrderById(response, request)

			return
		}

		http.Error(response, "несуществующий URL", http.StatusNotFound)
	default:
		http.Error(response, "несуществующий http метод", http.StatusMethodNotAllowed)
	}
}

func (a *api) GetOrderById(response http.ResponseWriter, request *http.Request) {
	segments := strings.Split(request.URL.Path, "/")

	if len(segments) != 4 || segments[1] != "v1" || segments[2] != "orders" {
		http.Error(response, "неверный URL", http.StatusBadRequest)

		return
	}

	id := segments[3]

	orderService := service.NewService()

	orderDto, err := orderService.GetOrderById(id)

	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			http.Error(response, fmt.Sprintf("пользователь не найден с order_uid: %s", id), http.StatusBadRequest)

			return
		}

		http.Error(response, err.Error(), http.StatusBadRequest)

		return
	}

	orderJson, err := json.Marshal(orderDto)

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)

		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	_, err = response.Write(orderJson)

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)

		return
	}
}
