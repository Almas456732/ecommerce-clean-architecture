package repositories

import (
	"context"
	"inventory_service/internal/domain"
	"inventory_service/internal/dto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CategoryRepositoryImpl struct {
	collection *mongo.Collection
}

func NewCategoryRepository(databaseURL string) *CategoryRepositoryImpl {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(databaseURL))
	if err != nil {
		panic(err)
	}

	collection := client.Database("inventory").Collection("categories")
	return &CategoryRepositoryImpl{collection: collection}
}

func (r *CategoryRepositoryImpl) Create(category *domain.Category) error {
	category.ID = primitive.NewObjectID().Hex()

	categoryDTO := &dto.CategoryDTO{}
	categoryDTO.FromDomain(category)

	_, err := r.collection.InsertOne(context.Background(), categoryDTO)
	return err
}

func (r *CategoryRepositoryImpl) FindByID(id string) (*domain.Category, error) {
	var categoryDTO dto.CategoryDTO
	err := r.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&categoryDTO)
	if err != nil {
		return nil, err
	}
	return categoryDTO.ToDomain(), nil
}

func (r *CategoryRepositoryImpl) Update(category *domain.Category) error {
	categoryDTO := &dto.CategoryDTO{}
	categoryDTO.FromDomain(category)

	_, err := r.collection.UpdateOne(
		context.Background(),
		bson.M{"_id": category.ID},
		bson.M{"$set": bson.M{
			"name":        category.Name,
			"description": category.Description,
		}},
	)
	return err
}

func (r *CategoryRepositoryImpl) Delete(id string) error {
	_, err := r.collection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}

func (r *CategoryRepositoryImpl) FindAll() ([]*domain.Category, error) {
	cursor, err := r.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	var categoryDTOs []dto.CategoryDTO
	if err = cursor.All(context.Background(), &categoryDTOs); err != nil {
		return nil, err
	}

	categories := make([]*domain.Category, len(categoryDTOs))
	for i, dto := range categoryDTOs {
		categories[i] = dto.ToDomain()
	}

	return categories, nil
}
