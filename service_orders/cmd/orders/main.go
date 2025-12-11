package main

import (
	"log"
	"net/http"
	"os"

	"pr2/service_orders/internal/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/v1/orders", handlers.CreateOrder).Methods("POST")
	r.HandleFunc("/v1/orders/{id}", handlers.GetOrder).Methods("GET")
	r.HandleFunc("/v1/orders", handlers.ListOrders).Methods("GET")
	r.HandleFunc("/v1/orders/{id}", handlers.UpdateOrder).Methods("PUT")
	r.HandleFunc("/v1/orders/{id}", handlers.CancelOrder).Methods("DELETE")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8002"
	}

	log.Println("Order Service running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
