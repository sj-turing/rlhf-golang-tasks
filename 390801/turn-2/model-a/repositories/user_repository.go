// repositories/user_repository.go
package repositories

import "390801/turn-2/model-a/models"

type UserRepository interface {
	Create(user *models.User) error
	GetByID(id string) (*models.User, error)
	/*
		Update(user *models.User) error
		Delete(id string) error
		List() ([]*models.User, error)
	*/
}
