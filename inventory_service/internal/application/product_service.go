package application

import (
	"inventory_service/internal/domain"
)

type ProductService struct {
	repo domain.ProductRepository
}

func NewProductService(repo domain.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(product *domain.Product) error {
	return s.repo.Create(product)
}

func (s *ProductService) GetProduct(id string) (*domain.Product, error) {
	return s.repo.FindByID(id)
}

func (s *ProductService) UpdateProduct(product *domain.Product) error {
	return s.repo.Update(product)
}

func (s *ProductService) DeleteProduct(id string) error {
	return s.repo.Delete(id)
}

func (s *ProductService) ListProducts(filter domain.ProductFilter, page, limit int) ([]*domain.Product, error) {
	return s.repo.FindAll(filter, page, limit)
}
