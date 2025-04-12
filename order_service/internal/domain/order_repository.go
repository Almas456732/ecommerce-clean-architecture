package domain

type OrderRepository interface {
	Create(order *Order) error
	FindByID(id string) (*Order, error)
	Update(order *Order) error
	FindByUserID(userID string) ([]*Order, error)
}
