package models

type Order struct {
	ID        string  `json:"id"`
	UserID    string  `json:"user_id"`
	Item      string  `json:"item"`
	Quantity  int     `json:"quantity"`
	Status    string  `json:"status"` // new, paid, shipped, cancelled
	Price     float64 `json:"price"`
	CreatedAt int64   `json:"created_at"`
}
