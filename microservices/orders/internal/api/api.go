package api

import (
	"net/http"
)

type OrdersApiInterface interface {
	OrdersHandler(http.ResponseWriter, *http.Request)
	GetOrderById(http.ResponseWriter, *http.Request)
}
