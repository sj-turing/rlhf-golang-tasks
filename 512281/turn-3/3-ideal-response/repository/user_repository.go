// For example, in repository/user_repository.go
package repository

import "512281/turn-3/3-ideal-response/models"

//go:generate mockgen -source user_repository.go -destination ../mocks/user_repository_mock.go -package mocks
type UserRepository interface {
	GetUserByUsername(username string) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(username string) error
}
