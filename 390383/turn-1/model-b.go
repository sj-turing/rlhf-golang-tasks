package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// APIClient interface defines the operations of the API client
type APIClient interface {
	GetUsers(ctx context.Context) ([]User, error)
}

// User struct represents a user
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// realAPIClient implements the APIClient interface
type realAPIClient struct {
	httpClient *http.Client
}

// NewRealAPIClient creates a new instance of realAPIClient
func NewRealAPIClient(httpClient *http.Client) APIClient {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	return &realAPIClient{httpClient}
}

// GetUsers fetches users from the API
func (c *realAPIClient) GetUsers(ctx context.Context) ([]User, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "https://api.example.com/users", nil)
	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var users []User
	if err := json.Unmarshal(body, &users); err != nil {
		return nil, err
	}

	return users, nil
}

// mockHTTPClient implements http.Client for testing purposes
type mockHTTPClient struct {
	mock.Mock
}

// Do implements the Do method of http.Client
func (m *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	args := m.Called(req)
	resp, ok := args.Get(0).(*http.Response)
	err := args.Error(1)

	if !ok {
		panic("mockHTTPClient.Do: expected *http.Response as first argument")
	}
	return resp, err
}

func TestGetUsers(t *testing.T) {
	// Create a mock HTTP client
	mockClient := &mockHTTPClient{}

	// Set up a mock response
	mockResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(bytes.NewReader([]byte(`[{"id": 1, "name": "Alice"}]`))),
	}

	// Configure the mock to return the mock response
	mockClient.On("Do", mock.Anything).Return(mockResponse, nil)

	// Create an API client with the mock HTTP client
	apiClient := NewRealAPIClient(mockClient)

	// Test GetUsers method
	users, err := apiClient.GetUsers(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, 1, len(users))
	assert.Equal(t, "Alice", users[0].Name)
}
