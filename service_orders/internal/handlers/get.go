package handlers

import (
	"encoding/json"
	"net/http"

	"pr2/service_orders/internal/service"

	"github.com/gorilla/mux"
)

func GetOrder(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("X-User-ID")
	if userID == "" {
		http.Error(w, "missing user id", http.StatusUnauthorized)
		return
	}

	id := mux.Vars(r)["id"]
	order, err := service.Get(id, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"success": true,
		"data":    order,
	})
}

func ListOrders(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("X-User-ID")
	if userID == "" {
		http.Error(w, "missing user id", http.StatusUnauthorized)
		return
	}

	orders := service.List(userID)

	json.NewEncoder(w).Encode(map[string]any{
		"success": true,
		"data":    orders,
	})
}
