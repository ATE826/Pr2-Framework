package handlers

import (
	"encoding/json"
	"net/http"

	"pr2/service_users/internal/service"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	user, err := service.RegisterUser(body.Email, body.Password, body.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "data": user})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	token, err := service.LoginUser(creds.Email, creds.Password)
	if err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    map[string]string{"token": token},
	})
}
