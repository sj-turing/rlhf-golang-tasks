package services_test

import (
	"context"
	"testing"

	"github.com/example/app/models"
	"github.com/example/app/repositories"
	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

func TestUserService_CreateUser(t *testing.T) {
	// Create a mock repository
	mockRepo := &mockRepositories.UserRepository{}

	// Define expectations
	mockRepo.On("Create", &models.User{ID: 1, Name: "John"}).Return(nil)

	// Create the user service with the mock repository
	userService := services.NewUserService(mockRepo)

	// Test the service
	err := userService.CreateUser(&models.User{ID: 1, Name: "John"})
	assert.NoError(t, err)

	// Verify the expectation
	mockRepo.AssertExpectations(t)
}

func TestUserService_GetUserByID(t *testing.T) {
	mockRepo := &mockRepositories.UserRepository{}
	expectedUser := &models.User{ID: 1, Name: "John"}

	mockRepo.On("GetByID", 1).Return(expectedUser, nil)

	userService := services.NewUserService(mockRepo)

	user, err := userService.GetUserByID(1)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)

	mockRepo.AssertExpectations(t)
}

// Add more tests for UpdateUser, DeleteUser, and ListUsers similarly
