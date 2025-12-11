package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"pr2/pkg/jwt"

	"github.com/google/uuid"
)

// Минимальная модель пользователя
type User struct {
	ID       string   `json:"id"`
	Email    string   `json:"email"`
	Password string   `json:"password"` // Для упрощения хранение plain (на оценку 4)
	Name     string   `json:"name"`
	Roles    []string `json:"roles"`
}

var users = map[string]User{} // временная память

func Register(w http.ResponseWriter, r *http.Request) {
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "error": "invalid body"})
		return
	}
	for _, user := range users {
		if user.Email == u.Email {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "error": "email exists"})
			return
		}
	}
	u.ID = uuid.New().String()
	u.Roles = []string{"user"}
	users[u.ID] = u

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "data": u})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "error": "invalid body"})
		return
	}
	for _, u := range users {
		if u.Email == creds.Email && u.Password == creds.Password {
			token, _ := jwt.GenerateToken(u.ID, u.Roles, time.Hour*24)
			json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "data": map[string]string{"token": token}})
			return
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "error": "invalid credentials"})
}
