package middleware

import (
	"log"
	"net/http"
	"os"
)

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		header := os.Getenv("AUTHENTICATION_TOKEN_HEADER")

		if header == "" {
			log.Panicf("specify the name of the authentication token")
		}

		token := os.Getenv("AUTHENTICATION_TOKEN")

		if token == "" {
			log.Panicf("specify the value of the authentication token")
		}

		key := request.Header.Get(header)

		if key != token {
			http.Error(response, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(response, request)
	})
}
