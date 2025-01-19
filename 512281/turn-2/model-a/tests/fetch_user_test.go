package service_test

import (
	"512281/turn-2/model-a/mocks"
	"512281/turn-2/model-a/models"
	"512281/turn-2/model-a/service"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestFetchUser(t *testing.T) {
	// Initialize gomock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock instance of UserRepository
	userRepoMock := mocks.NewMockUserRepository(ctrl)

	// Set up expectations
	userRepoMock.
		EXPECT().
		GetUserByUsername("testuser").
		Return(&models.User{Username: "testuser"}, nil).
		Times(1)

	// Initialize the service with the mock repository
	userService := service.UserService{Repo: userRepoMock}

	// Perform the test
	user, err := userService.FetchUser("testuser")

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, "testuser", user.Username)
}
