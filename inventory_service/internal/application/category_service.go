package application

import (
	"inventory_service/internal/domain"
)

type CategoryService struct {
	repo domain.CategoryRepository
}

func NewCategoryService(repo domain.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) CreateCategory(category *domain.Category) error {
	return s.repo.Create(category)
}

func (s *CategoryService) GetCategory(id string) (*domain.Category, error) {
	return s.repo.FindByID(id)
}

func (s *CategoryService) UpdateCategory(category *domain.Category) error {
	return s.repo.Update(category)
}

func (s *CategoryService) DeleteCategory(id string) error {
	return s.repo.Delete(id)
}

func (s *CategoryService) ListCategories() ([]*domain.Category, error) {
	return s.repo.FindAll()
}
