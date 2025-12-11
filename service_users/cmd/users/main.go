package main

import (
	"log"
	"net/http"

	"pr2/service_users/internal/config"
	"pr2/service_users/internal/handlers"

	"github.com/gorilla/mux"
)

func main() {

	cfg := config.Load()

	r := mux.NewRouter()

	// Handlers
	r.HandleFunc("/v1/users/register", handlers.Register).Methods("POST")
	r.HandleFunc("/v1/users/login", handlers.Login).Methods("POST")
	r.HandleFunc("/v1/users/profile", handlers.Profile).Methods("GET")

	log.Println("User service running on port", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}
