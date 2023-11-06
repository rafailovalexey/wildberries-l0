package http_server

import (
	"fmt"
	"github.com/emptyhopes/orders/cmd/http-server/interceptor"
	"github.com/emptyhopes/orders/cmd/http-server/middleware"
	"github.com/emptyhopes/orders/internal/api/orders"
	"log"
	"net/http"
	"os"
)

func Run() {
	router := http.NewServeMux()

	middlewares := middleware.ChainMiddleware(
		interceptor.LoggingInterceptor,
		middleware.AuthenticationMiddleware,
	)

	Orders := &orders.Api{}

	router.Handle("/v1/orders/", middlewares(http.HandlerFunc(Orders.OrdersHandler)))

	hostname := os.Getenv("HOSTNAME")

	if hostname == "" {
		log.Fatalf("укажите имя хоста")
	}

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatalf("укажите порт")
	}

	address := fmt.Sprintf("%s:%s", hostname, port)

	fmt.Println(fmt.Sprintf("сервер запускается по адресу %s", address))

	err := http.ListenAndServe(address, router)

	if err != nil {
		log.Fatalf("ошибка при запуске сервера: %v\n", err)
	}
}
