package repositories

import (
	"context"
	"errors"
	"time"
	"user_service/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoUserRepository implements UserRepository with MongoDB
type MongoUserRepository struct {
	collection *mongo.Collection
}

// MongoUser is a MongoDB representation of a User
type MongoUser struct {
	ID        string    `bson:"_id"`
	Username  string    `bson:"username"`
	Password  string    `bson:"password"`
	Email     string    `bson:"email"`
	FullName  string    `bson:"full_name"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

// NewMongoUserRepository creates a new MongoDB user repository
func NewMongoUserRepository(uri, dbName, collectionName string) (*MongoUserRepository, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	// Ping the MongoDB server to check connection
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	collection := client.Database(dbName).Collection(collectionName)

	// Create indexes for fast lookups by username and email
	_, err = collection.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys:    bson.D{bson.E{Key: "username", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bson.D{bson.E{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	})
	if err != nil {
		return nil, err
	}

	return &MongoUserRepository{
		collection: collection,
	}, nil
}

// toDomainUser converts a MongoDB user to a domain User
func toDomainUser(mu *MongoUser) *domain.User {
	return &domain.User{
		ID:        mu.ID,
		Username:  mu.Username,
		Password:  mu.Password,
		Email:     mu.Email,
		FullName:  mu.FullName,
		CreatedAt: mu.CreatedAt,
		UpdatedAt: mu.UpdatedAt,
	}
}

// toMongoUser converts a domain User to a MongoDB user
func toMongoUser(u *domain.User) *MongoUser {
	return &MongoUser{
		ID:        u.ID,
		Username:  u.Username,
		Password:  u.Password,
		Email:     u.Email,
		FullName:  u.FullName,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

// Save adds a new user to the repository
func (r *MongoUserRepository) Save(user *domain.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, toMongoUser(user))
	if mongo.IsDuplicateKeyError(err) {
		return errors.New("user with this username or email already exists")
	}
	return err
}

// FindByID retrieves a user by ID
func (r *MongoUserRepository) FindByID(id string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var result MongoUser
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return toDomainUser(&result), nil
}

// FindByUsername retrieves a user by username
func (r *MongoUserRepository) FindByUsername(username string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var result MongoUser
	err := r.collection.FindOne(ctx, bson.M{"username": username}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return toDomainUser(&result), nil
}

// FindByEmail retrieves a user by email
func (r *MongoUserRepository) FindByEmail(email string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var result MongoUser
	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return toDomainUser(&result), nil
}

// Update updates an existing user
func (r *MongoUserRepository) Update(user *domain.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": toMongoUser(user)}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return errors.New("email already in use")
		}
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("user not found")
	}

	return nil
}

// Delete removes a user by ID
func (r *MongoUserRepository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("user not found")
	}

	return nil
}

// List returns all users
func (r *MongoUserRepository) List() ([]*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var mongoUsers []MongoUser
	if err := cursor.All(ctx, &mongoUsers); err != nil {
		return nil, err
	}

	users := make([]*domain.User, len(mongoUsers))
	for i, mu := range mongoUsers {
		users[i] = toDomainUser(&mu)
	}

	return users, nil
}
