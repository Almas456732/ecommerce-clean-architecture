package application

import (
	"order_service/internal/domain"
	"time"
)

type OrderService struct {
	repo domain.OrderRepository
}

func NewOrderService(repo domain.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(order *domain.Order) error {
	
	var total float64
	for _, item := range order.Items {
		total += item.Price * float64(item.Quantity)
	}
	order.Total = total
	order.Status = domain.OrderStatusPending
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()

	return s.repo.Create(order)
}

func (s *OrderService) GetOrder(id string) (*domain.Order, error) {
	return s.repo.FindByID(id)
}

func (s *OrderService) UpdateOrderStatus(order *domain.Order) error {
	order.UpdatedAt = time.Now()
	return s.repo.Update(order)
}

func (s *OrderService) GetUserOrders(userID string) ([]*domain.Order, error) {
	return s.repo.FindByUserID(userID)
}
