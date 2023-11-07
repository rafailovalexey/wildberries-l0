package api

import (
	"net/http"
)

type OrderApiInterface interface {
	OrdersHandler(http.ResponseWriter, *http.Request)
	GetOrderById(http.ResponseWriter, *http.Request)
}
