package db

import "pr2/service_users/internal/models"

var Users = map[string]models.User{}

func SaveUser(u models.User) {
	Users[u.ID] = u
}

func FindByEmail(email string) *models.User {
	for _, u := range Users {
		if u.Email == email {
			return &u
		}
	}
	return nil
}

func FindByID(id string) *models.User {
	if u, ok := Users[id]; ok {
		return &u
	}
	return nil
}
