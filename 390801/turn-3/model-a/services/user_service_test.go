// services/user_service_test.go
package services_test

import (
	"testing"

	"390801/turn-3/model-a/models"
	"390801/turn-3/model-a/services"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository is a mock type for the UserRepository interface
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) GetByID(id string) (*models.User, error) {
	args := m.Called(id)
	return args.Get(0).(*models.User), args.Error(1)
}

// Test_CreateUser tests the CreateUser method of UserService
func Test_CreateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userService := services.NewUserService(mockRepo)
	user := &models.User{ID: "1", Name: "Alice"}

	// Define the behavior of the mock
	mockRepo.On("Create", user).Return(nil)

	err := userService.CreateUser(user)

	// Assertions
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

// Test_GetUserByID tests the GetUserByID method of UserService
func Test_GetUserByID(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userService := services.NewUserService(mockRepo)
	user := &models.User{ID: "1", Name: "Alice"}

	// Define the behavior of the mock
	mockRepo.On("GetByID", "1").Return(user, nil)

	result, err := userService.GetUserByID("1")

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, user, result)
	mockRepo.AssertExpectations(t)
}

// Add additional test cases to handle errors or special conditions as needed
