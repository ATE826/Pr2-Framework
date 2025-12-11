package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"pr2/api_gateway/internal/middleware"
	"pr2/api_gateway/internal/proxy"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.CORS)

	// JWT только для защищенных путей
	r.Use(middleware.AuthMiddleware)

	// Прокси маршрутов
	r.PathPrefix("/v1/users").Handler(proxy.UsersProxy())
	r.PathPrefix("/v1/orders").Handler(proxy.OrdersProxy())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Handler:      r,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("API Gateway running on port", port)
	log.Fatal(srv.ListenAndServe())
}
