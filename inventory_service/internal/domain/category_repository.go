package domain

// CategoryRepository defines operations for interacting with category storage
type CategoryRepository interface {
	Create(category *Category) error
	FindByID(id string) (*Category, error)
	Update(category *Category) error
	Delete(id string) error
	FindAll() ([]*Category, error)
}
