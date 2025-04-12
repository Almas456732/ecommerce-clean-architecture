package dto

import (
	"time"

	"inventory_service/internal/domain"
)

type ProductDTO struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	Name        string    `json:"name" bson:"name"`
	Description string    `json:"description" bson:"description"`
	CategoryID  string    `json:"category_id" bson:"category_id"`
	Price       float64   `json:"price" bson:"price"`
	Stock       int       `json:"quantity" bson:"stock"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}

func (dto *ProductDTO) FromDomain(product *domain.Product) {
	dto.ID = product.ID
	dto.Name = product.Name
	dto.Description = product.Description
	dto.CategoryID = product.CategoryID
	dto.Price = product.Price
	dto.Stock = product.Stock
	dto.CreatedAt = product.CreatedAt
	dto.UpdatedAt = product.UpdatedAt
}

func (dto *ProductDTO) ToDomain() *domain.Product {
	return &domain.Product{
		ID:          dto.ID,
		Name:        dto.Name,
		Description: dto.Description,
		CategoryID:  dto.CategoryID,
		Price:       dto.Price,
		Stock:       dto.Stock,
		CreatedAt:   dto.CreatedAt,
		UpdatedAt:   dto.UpdatedAt,
	}
}
