// For example, in repository/user_repository.go
package repository

import "512281/turn-3/model-a/models"

//go:generate mockgen -destination=../mocks/user_repository_mock.go -package=mocks myapp/repository UserRepository

type UserRepository interface {
	GetUserByUsername(username string) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(username string) error
}
