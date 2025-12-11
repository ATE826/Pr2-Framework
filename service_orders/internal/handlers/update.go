package handlers

import (
	"encoding/json"
	"net/http"

	"pr2/service_orders/internal/service"

	"github.com/gorilla/mux"
)

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("X-User-ID")
	if userID == "" {
		http.Error(w, "missing user id", http.StatusUnauthorized)
		return
	}

	id := mux.Vars(r)["id"]

	var body struct {
		Status string `json:"status"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	order, err := service.Update(id, userID, body.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"success": true,
		"data":    order,
	})
}
