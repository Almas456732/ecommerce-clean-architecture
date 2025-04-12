package repositories

import (
	"context"
	"inventory_service/internal/domain"
	"inventory_service/internal/dto"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProductRepositoryImpl struct {
	collection *mongo.Collection
}

func NewProductRepository(databaseURL string) *ProductRepositoryImpl {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(databaseURL))
	if err != nil {
		panic(err)
	}

	collection := client.Database("inventory").Collection("products")
	return &ProductRepositoryImpl{collection: collection}
}

func (r *ProductRepositoryImpl) Create(product *domain.Product) error {
	product.ID = primitive.NewObjectID().Hex()
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	productDTO := &dto.ProductDTO{}
	productDTO.FromDomain(product)

	_, err := r.collection.InsertOne(context.Background(), productDTO)
	return err
}

func (r *ProductRepositoryImpl) FindByID(id string) (*domain.Product, error) {
	var productDTO dto.ProductDTO
	err := r.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&productDTO)
	if err != nil {
		return nil, err
	}
	return productDTO.ToDomain(), nil
}

func (r *ProductRepositoryImpl) Update(product *domain.Product) error {
	product.UpdatedAt = time.Now()

	productDTO := &dto.ProductDTO{}
	productDTO.FromDomain(product)

	_, err := r.collection.UpdateOne(
		context.Background(),
		bson.M{"_id": product.ID},
		bson.M{"$set": bson.M{
			"name":        product.Name,
			"description": product.Description,
			"category_id": product.CategoryID,
			"price":       product.Price,
			"stock":       product.Stock,
			"updated_at":  product.UpdatedAt,
		}},
	)
	return err
}

func (r *ProductRepositoryImpl) Delete(id string) error {
	_, err := r.collection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}

func (r *ProductRepositoryImpl) FindAll(filter domain.ProductFilter, page, limit int) ([]*domain.Product, error) {
	query := bson.M{}

	if filter.Name != "" {
		query["name"] = bson.M{"$regex": filter.Name, "$options": "i"}
	}
	if filter.CategoryID != "" {
		query["category_id"] = filter.CategoryID
	}
	if filter.MinPrice > 0 {
		query["price"] = bson.M{"$gte": filter.MinPrice}
	}
	if filter.MaxPrice > 0 {
		if query["price"] == nil {
			query["price"] = bson.M{"$lte": filter.MaxPrice}
		} else {
			query["price"].(bson.M)["$lte"] = filter.MaxPrice
		}
	}

	opts := options.Find().
		SetSkip(int64((page - 1) * limit)).
		SetLimit(int64(limit))

	cursor, err := r.collection.Find(context.Background(), query, opts)
	if err != nil {
		return nil, err
	}

	var productDTOs []dto.ProductDTO
	if err = cursor.All(context.Background(), &productDTOs); err != nil {
		return nil, err
	}

	products := make([]*domain.Product, len(productDTOs))
	for i, dto := range productDTOs {
		products[i] = dto.ToDomain()
	}

	return products, nil
}
