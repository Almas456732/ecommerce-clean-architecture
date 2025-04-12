package dto

import (
	"time"

	"order_service/internal/domain"
)

type OrderItemDTO struct {
	ProductID string  `json:"product_id" bson:"product_id"`
	Quantity  int     `json:"quantity" bson:"quantity"`
	Price     float64 `json:"price" bson:"price"`
}

type OrderDTO struct {
	ID        string         `json:"id" bson:"_id,omitempty"`
	UserID    string         `json:"user_id" bson:"user_id"`
	Items     []OrderItemDTO `json:"items" bson:"items"`
	Total     float64        `json:"total" bson:"total"`
	Status    string         `json:"status" bson:"status"`
	CreatedAt time.Time      `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" bson:"updated_at"`
}

func (dto *OrderDTO) FromDomain(order *domain.Order) {
	dto.ID = order.ID
	dto.UserID = order.UserID
	dto.Total = order.Total
	dto.Status = string(order.Status)
	dto.CreatedAt = order.CreatedAt
	dto.UpdatedAt = order.UpdatedAt

	dto.Items = make([]OrderItemDTO, len(order.Items))
	for i, item := range order.Items {
		dto.Items[i] = OrderItemDTO{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Price,
		}
	}
}

func (dto *OrderDTO) ToDomain() *domain.Order {
	items := make([]domain.OrderItem, len(dto.Items))
	for i, item := range dto.Items {
		items[i] = domain.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Price,
		}
	}

	return &domain.Order{
		ID:        dto.ID,
		UserID:    dto.UserID,
		Items:     items,
		Total:     dto.Total,
		Status:    domain.OrderStatus(dto.Status),
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
	}
}
