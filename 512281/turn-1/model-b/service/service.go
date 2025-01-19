// service.go
package service

import (
	"512281/turn-1/model-a/models"
	"512281/turn-1/model-a/repository"
	"context"

	"github.com/pkg/errors"
)

type UserService interface {
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) error
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id int) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

// Implement UserService methods using UserRepository interface
func (s *userService) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	if username == "" {
		return nil, errors.New("username is required")
	}
	return s.repo.GetUserByUsername(ctx, username)
}

// ... (Other service methods)
