package db

import "pr2/service_orders/internal/models"

var Orders = map[string]models.Order{}

func SaveOrder(o models.Order) {
	Orders[o.ID] = o
}

func GetOrder(id string) *models.Order {
	if o, ok := Orders[id]; ok {
		return &o
	}
	return nil
}

func ListByUser(userID string) []models.Order {
	var result []models.Order
	for _, o := range Orders {
		if o.UserID == userID {
			result = append(result, o)
		}
	}
	return result
}
