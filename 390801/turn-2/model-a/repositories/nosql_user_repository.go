// repositories/nosql_user_repository.go
package repositories

import (
	"390801/turn-2/model-a/models"
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type NoSQLUserRepository struct {
	collection *mongo.Collection
}

func NewNoSQLUserRepository(client *mongo.Client, dbName string) UserRepository {
	collection := client.Database(dbName).Collection("users")
	return &NoSQLUserRepository{collection: collection}
}

func (r *NoSQLUserRepository) Create(user *models.User) error {
	_, err := r.collection.InsertOne(context.Background(), user)
	return err
}

func (r *NoSQLUserRepository) GetByID(id string) (*models.User, error) {
	user := &models.User{}
	err := r.collection.FindOne(context.Background(), bson.M{"id": id}).Decode(user)
	if err != nil {
		return nil, errors.Wrap(err, "user not found")
	}
	return user, nil
}

// Similar implementations for Update, Delete, and List
