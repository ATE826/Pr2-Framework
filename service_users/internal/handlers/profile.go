package handlers

import (
	"encoding/json"
	"net/http"

	"pr2/service_users/internal/service"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("X-User-ID")
	if userID == "" {
		http.Error(w, "missing user id", http.StatusUnauthorized)
		return
	}

	user, err := service.GetProfile(userID)
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    user,
	})
}
