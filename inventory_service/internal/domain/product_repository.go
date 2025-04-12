package domain

type ProductRepository interface {
	Create(product *Product) error
	FindByID(id string) (*Product, error)
	Update(product *Product) error
	Delete(id string) error
	FindAll(filter ProductFilter, page, limit int) ([]*Product, error)
}
