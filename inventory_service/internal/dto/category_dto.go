package dto

import "inventory_service/internal/domain"

type CategoryDTO struct {
	ID          string `json:"id" bson:"_id,omitempty"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
}

func (dto *CategoryDTO) FromDomain(category *domain.Category) {
	dto.ID = category.ID
	dto.Name = category.Name
	dto.Description = category.Description
}

func (dto *CategoryDTO) ToDomain() *domain.Category {
	return &domain.Category{
		ID:          dto.ID,
		Name:        dto.Name,
		Description: dto.Description,
	}
}
