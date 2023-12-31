package orders

import (
	"encoding/json"
	"fmt"
	definition "github.com/emptyhopes/orders/internal/api"
	"github.com/emptyhopes/orders/internal/service"
	"github.com/emptyhopes/orders/internal/validation"
	"net/http"
	"strings"
)

type api struct {
	orderValidation validation.OrderValidationInterface
	orderService    service.OrderServiceInterface
}

var _ definition.OrderApiInterface = (*api)(nil)

func NewOrderApi(orderValidation validation.OrderValidationInterface, orderService service.OrderServiceInterface) *api {
	return &api{
		orderValidation: orderValidation,
		orderService:    orderService,
	}
}

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

		http.Error(response, getErrorJson("несуществующий URL"), http.StatusNotFound)
	default:
		http.Error(response, getErrorJson("несуществующий http метод"), http.StatusMethodNotAllowed)
	}
}

func (a *api) GetOrderById(response http.ResponseWriter, request *http.Request) {
	segments := strings.Split(request.URL.Path, "/")

	if len(segments) != 4 || segments[1] != "v1" || segments[2] != "orders" {
		http.Error(response, getErrorJson("неверный URL"), http.StatusBadRequest)

		return
	}

	id := segments[3]

	err := a.orderValidation.GetOrderByIdValidation(id)

	if err != nil {
		http.Error(response, getErrorJson(err.Error()), http.StatusBadRequest)

		return
	}

	orderDto, err := a.orderService.GetOrderById(id)

	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			http.Error(response, getErrorJson(fmt.Sprintf("заказ не найден с order_uid: %s", id)), http.StatusBadRequest)

			return
		}

		http.Error(response, getErrorJson(err.Error()), http.StatusBadRequest)

		return
	}

	orderJson, err := json.Marshal(orderDto)

	if err != nil {
		http.Error(response, getErrorJson(err.Error()), http.StatusInternalServerError)

		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	_, err = response.Write(orderJson)

	if err != nil {
		http.Error(response, getErrorJson(err.Error()), http.StatusInternalServerError)

		return
	}
}

func getErrorJson(message string) string {
	return fmt.Sprintf("{\"error\":\"%s\"}", message)
}
