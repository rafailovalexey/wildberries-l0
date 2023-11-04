package orders

import (
	"fmt"
	"github.com/emptyhopes/level0/internal/api"
	service "github.com/emptyhopes/level0/internal/service/orders"
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

	fmt.Println(id)

	orderService := &service.Service{}

	result, err := orderService.GetOrderById(id)

	if err != nil {
		http.Error(response, "ошибка при получение заказа", http.StatusBadRequest)
		return
	}

	fmt.Println(result)
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
