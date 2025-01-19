package mocks

import (
	"512281/turn-1/model-a/models"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) GetUserByUsername(username string) (*models.User, error) {
	args := m.Called(username)
	return args.Get(0).(*models.User), args.Error(1)
}

// Implement mock methods for CreateUser, UpdateUser, DeleteUser
func (m *UserRepositoryMock) CreateUser(user *models.User) error {
	return nil
}

func (m *UserRepositoryMock) UpdateUser(user *models.User) error {
	return nil
}

func (m *UserRepositoryMock) DeleteUser(username string) error {
	return nil
}
