package repository

import "512281/turn-2/model-b/models"

type UserRepository interface {
	GetUserByUsername(username string) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(username string) error
}
