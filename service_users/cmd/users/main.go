package main

import (
	"log"
	"net/http"
	"os"

	"pr2/service_users/internal/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Auth routes
	r.HandleFunc("/v1/users/register", handlers.Register).Methods("POST")
	r.HandleFunc("/v1/users/login", handlers.Login).Methods("POST")
	r.HandleFunc("/v1/users/profile", handlers.Profile).Methods("GET", "PUT")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8001"
	}

	log.Println("User Service running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
