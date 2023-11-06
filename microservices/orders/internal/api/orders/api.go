package orders

import (
	"encoding/json"
	"github.com/emptyhopes/orders/internal/api"
	service "github.com/emptyhopes/orders/internal/service/orders"
	"net/http"
	"strings"
)

type Api struct{}

var _ api.OrdersApiInterface = &Api{}

func (a *Api) GetOrderById(response http.ResponseWriter, request *http.Request) {
	segments := strings.Split(request.URL.Path, "/")

	if len(segments) != 4 || segments[1] != "v1" || segments[2] != "orders" {
		http.Error(response, "неверный URL", http.StatusBadRequest)
		return
	}

	id := segments[3]

	orderService := &service.Service{}

	orderDto, err := orderService.GetOrderById(id)

	if err != nil {
		http.Error(response, "ошибка при получение заказа", http.StatusBadRequest)
		return
	}

	orderJson, err := json.Marshal(orderDto)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(orderJson)
}

/*
Здесь я использовал парсинг URL, для того, чтобы добиться REST поведения
GetAllOrders - /v1/orders
GetOrderById - /v1/orders/:id
*/
func (a *Api) OrdersHandler(response http.ResponseWriter, request *http.Request) {
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
