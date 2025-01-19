// repository.go
package repository

import (
	"512281/turn-1/model-a/models"
	"context"
)

type UserRepository interface {
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) error
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id int) error
}
