// repositories/user_repository.go
package repositories

import "390801/turn-3/model-a/models"

type UserRepository interface {
	Create(user *models.User) error
	GetByID(id string) (*models.User, error)
	// Other repository methods...
}
