package domain

import "time"

type Product struct {
	ID          string
	Name        string
	Description string
	CategoryID  string
	Price       float64
	Stock       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ProductFilter struct {
	Name       string
	CategoryID string
	MinPrice   float64
	MaxPrice   float64
}
