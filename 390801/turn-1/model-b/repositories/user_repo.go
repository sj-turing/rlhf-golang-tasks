package repositories

import "390801/turn-1/model-b/models"

type UserRepository interface {
	Create(user *models.User) error
	GetByID(id int) (*models.User, error)
	Update(user *models.User) error
	Delete(id int) error
	List() ([]*models.User, error)
}
