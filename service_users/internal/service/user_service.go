package service

import (
	"errors"
	"time"

	"pr2/pkg/jwt"
	"pr2/service_users/internal/db"
	"pr2/service_users/internal/models"

	"github.com/google/uuid"
)

func RegisterUser(email, password, name string) (*models.User, error) {
	if db.FindByEmail(email) != nil {
		return nil, errors.New("email exists")
	}

	user := models.User{
		ID:       uuid.New().String(),
		Email:    email,
		Password: password,
		Name:     name,
		Roles:    []string{"user"},
	}

	db.SaveUser(user)
	return &user, nil
}

func LoginUser(email, pass string) (string, error) {
	u := db.FindByEmail(email)
	if u == nil || u.Password != pass {
		return "", errors.New("invalid credentials")
	}

	token, _ := jwt.GenerateToken(u.ID, u.Roles, time.Hour*24)
	return token, nil
}

func GetProfile(userID string) (*models.User, error) {
	user := db.FindByID(userID)
	if user == nil {
		return nil, errors.New("not found")
	}
	return user, nil
}
