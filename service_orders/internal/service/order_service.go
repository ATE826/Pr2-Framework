package service

import (
	"errors"
	"time"

	"pr2/service_orders/internal/db"
	"pr2/service_orders/internal/models"

	"github.com/google/uuid"
)

func Create(userID, item string, qty int, price float64) *models.Order {
	order := models.Order{
		ID:        uuid.New().String(),
		UserID:    userID,
		Item:      item,
		Quantity:  qty,
		Price:     price,
		Status:    "new",
		CreatedAt: time.Now().Unix(),
	}
	db.SaveOrder(order)
	return &order
}

func Get(orderID, userID string) (*models.Order, error) {
	o := db.GetOrder(orderID)
	if o == nil {
		return nil, errors.New("not found")
	}
	if o.UserID != userID {
		return nil, errors.New("forbidden")
	}
	return o, nil
}

func List(userID string) []models.Order {
	return db.ListByUser(userID)
}

func Update(orderID, userID, status string) (*models.Order, error) {
	o := db.GetOrder(orderID)
	if o == nil {
		return nil, errors.New("not found")
	}
	if o.UserID != userID {
		return nil, errors.New("forbidden")
	}
	o.Status = status
	db.SaveOrder(*o)
	return o, nil
}
