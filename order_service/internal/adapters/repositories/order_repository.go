package repositories

import (
	"context"
	"order_service/internal/domain"
	"order_service/internal/dto"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OrderRepositoryImpl struct {
	collection *mongo.Collection
}

func NewOrderRepository(databaseURL string) *OrderRepositoryImpl {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(databaseURL))
	if err != nil {
		panic(err)
	}

	collection := client.Database("orders").Collection("orders")
	return &OrderRepositoryImpl{collection: collection}
}

func (r *OrderRepositoryImpl) Create(order *domain.Order) error {
	order.ID = primitive.NewObjectID().Hex()
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()

	if order.Total == 0 {
		for _, item := range order.Items {
			order.Total += item.Price * float64(item.Quantity)
		}
	}

	orderDTO := &dto.OrderDTO{}
	orderDTO.FromDomain(order)

	_, err := r.collection.InsertOne(context.Background(), orderDTO)
	return err
}

func (r *OrderRepositoryImpl) FindByID(id string) (*domain.Order, error) {
	var orderDTO dto.OrderDTO
	err := r.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&orderDTO)
	if err != nil {
		return nil, err
	}
	return orderDTO.ToDomain(), nil
}

func (r *OrderRepositoryImpl) Update(order *domain.Order) error {
	order.UpdatedAt = time.Now()

	orderDTO := &dto.OrderDTO{}
	orderDTO.FromDomain(order)

	_, err := r.collection.UpdateOne(
		context.Background(),
		bson.M{"_id": order.ID},
		bson.M{"$set": bson.M{
			"status":     order.Status,
			"updated_at": order.UpdatedAt,
		}},
	)
	return err
}

func (r *OrderRepositoryImpl) FindByUserID(userID string) ([]*domain.Order, error) {
	cursor, err := r.collection.Find(context.Background(), bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}

	var orderDTOs []dto.OrderDTO
	if err = cursor.All(context.Background(), &orderDTOs); err != nil {
		return nil, err
	}

	orders := make([]*domain.Order, len(orderDTOs))
	for i, dto := range orderDTOs {
		orders[i] = dto.ToDomain()
	}

	return orders, nil
}
