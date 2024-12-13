package services

import (
	"github.com/example/app/models"
	"github.com/example/app/repositories"
)

type UserService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.userRepo.Create(user)
}

func (s *UserService) GetUserByID(id int) (*models.User, error) {
	return s.userRepo.GetByID(id)
}

func (s *UserService) UpdateUser(user *models.User) error {
	return s.userRepo.Update(user)
}

func (s *UserService) DeleteUser(id int) error {
	return s.userRepo.Delete(id)
}

func (s *UserService) ListUsers() ([]*models.User, error) {
	return s.userRepo.List()
}
