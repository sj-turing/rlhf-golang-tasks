package service

import (
	"512281/turn-2/model-a/models"
	"512281/turn-2/model-a/repository"
)

type UserService struct {
	Repo repository.UserRepository
}

func (us *UserService) FetchUser(username string) (*models.User, error) {
	return us.Repo.GetUserByUsername(username)
}

// Implement Create, Update, and Delete using us.Repo.*
