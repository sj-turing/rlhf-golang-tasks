package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/example/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoUserRepository struct {
	client *mongo.Client
	db     *mongo.Database
}

func NewMongoUserRepository(url string, dbName string) (UserRepository, error) {
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return &mongoUserRepository{client: client, db: client.Database(dbName)}, nil
}

func (r *mongoUserRepository) Create(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := r.db.Collection("users").InsertOne(ctx, bson.D{{"id", user.ID}, {"name", user.Name}, {"email", user.Email}})
	return err
}

func (r *mongoUserRepository) GetByID(id int) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var user models.User
	err := r.db.Collection("users").FindOne(ctx, bson.D{{"id", id}}).Decode(&user)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

// ... similar implementations for Update, Delete, and List
