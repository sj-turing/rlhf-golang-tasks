package tests

import (
	"512281/turn-1/model-a/mocks"
	"512281/turn-1/model-a/models"
	"512281/turn-1/model-a/service"
	"testing"

	"github.com/stretchr/testify/assert"
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

