package domain

import (
	"time"
)

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusCompleted OrderStatus = "completed"
	OrderStatusCancelled OrderStatus = "cancelled"
)

type OrderItem struct {
	ProductID string
	Quantity  int
	Price     float64
}

type Order struct {
	ID        string
	UserID    string
	Items     []OrderItem
	Total     float64
	Status    OrderStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}
