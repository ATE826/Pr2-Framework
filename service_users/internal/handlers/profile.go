package handlers

import (
	"encoding/json"
	"net/http"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("X-User-ID") // для оценки 4 передаём из middleware
	u, ok := users[userID]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "error": "user not found"})
		return
	}
	if r.Method == "GET" {
		json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "data": u})
		return
	}
	if r.Method == "PUT" {
		var update User
		_ = json.NewDecoder(r.Body).Decode(&update)
		u.Name = update.Name
		users[userID] = u
		json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "data": u})
	}
}
