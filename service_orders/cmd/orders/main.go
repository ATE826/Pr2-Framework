package main

import (
	"log"
	"net/http"

	"pr2/service_orders/internal/config"
	"pr2/service_orders/internal/handlers"

	"github.com/gorilla/mux"
)

func main() {

	cfg := config.Load()

	r := mux.NewRouter()

	r.HandleFunc("/v1/orders", handlers.CreateOrder).Methods("POST")
	r.HandleFunc("/v1/orders", handlers.ListOrders).Methods("GET")
	r.HandleFunc("/v1/orders/{id}", handlers.GetOrder).Methods("GET")
	r.HandleFunc("/v1/orders/{id}", handlers.UpdateOrder).Methods("PUT")

	log.Println("Order service running on port", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}
