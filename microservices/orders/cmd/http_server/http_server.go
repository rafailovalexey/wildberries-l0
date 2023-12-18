package http_server

import (
	"fmt"
	"github.com/emptyhopes/orders/cmd/http_server/interceptor"
	"github.com/emptyhopes/orders/cmd/http_server/middleware"
	"github.com/emptyhopes/orders/internal/api"
	"log"
	"net/http"
	"os"
)

func Run(orderApi api.OrderApiInterface) {
	router := http.NewServeMux()

	middlewares := middleware.ChainMiddleware(
		interceptor.LoggingInterceptor,
		middleware.CorsMiddleware,
		middleware.AuthenticationMiddleware,
	)

	router.Handle("/v1/orders/", middlewares(http.HandlerFunc(orderApi.OrdersHandler)))

	hostname := os.Getenv("HOSTNAME")

	if hostname == "" {
		log.Panicf("enter the hostname")
	}

	port := os.Getenv("PORT")

	if port == "" {
		log.Panicf("specify the port")
	}

	address := fmt.Sprintf("%s:%s", hostname, port)

	log.Printf("the server starts at the address %s\n", address)

	err := http.ListenAndServe(address, router)

	if err != nil {
		log.Panicf("error when starting the server %v\n", err)
	}
}
