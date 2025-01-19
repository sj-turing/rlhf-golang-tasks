package tests

import (
	"github.com/stretchr/testify/assert"
	"myapp/mocks"
	"myapp/models"
	"myapp/service"
	"testing"
)

func TestFetchUser(t *testing.T) {
	userRepo := new(mocks.UserRepositoryMock)

	userRepo.On("GetUserByUsername", "testuser").Return(&models.User{Username: "testuser"}, nil)

	userService := service.UserService{Repo: userRepo}
	user, err := userService.FetchUser("testuser")

	assert.NoError(t, err)
	assert.Equal(t, "testuser", user.Username)
	userRepo.AssertExpectations(t)
}
