package handlers

import (
	"encoding/json"
	"net/http"

	"pr2/service_orders/internal/service"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("X-User-ID")
	if userID == "" {
		http.Error(w, "missing user id", http.StatusUnauthorized)
		return
	}

	var body struct {
		Item     string  `json:"item"`
		Quantity int     `json:"quantity"`
		Price    float64 `json:"price"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	order := service.Create(userID, body.Item, body.Quantity, body.Price)

	json.NewEncoder(w).Encode(map[string]any{
		"success": true,
		"data":    order,
	})
}
