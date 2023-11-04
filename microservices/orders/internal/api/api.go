package api

import (
	"net/http"
)

type OrdersApiInterface interface {
	GetOrderById(http.ResponseWriter, *http.Request)
	OrdersHandler(http.ResponseWriter, *http.Request)
}
