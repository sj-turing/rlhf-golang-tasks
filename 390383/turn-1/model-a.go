package main

/*
import (
	"encoding/json"
	"fmt"
	"net/http"
)

// User represents a user from the API
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// HTTPClient is an interface for making HTTP requests
type HTTPClient interface {
	Get(url string) (*http.Response, error)
}

// APIClient is our client for the JSON API
type APIClient struct {
	httpClient HTTPClient
}

// NewAPIClient creates a new APIClient with the provided HTTP client
func NewAPIClient(httpClient HTTPClient) *APIClient {
	return &APIClient{
		httpClient: httpClient,
	}
}

// GetUser fetches a user by ID from the JSON API
func (c *APIClient) GetUser(id int) (*User, error) {
	resp, err := c.httpClient.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%d", id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetching user: %s", resp.Status)
	}

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

// RealHTTPClient is a concrete implementation of HTTPClient
type RealHTTPClient struct{}

// Get performs a GET request
func (c *RealHTTPClient) Get(url string) (*http.Response, error) {
	return http.Get(url)
}

// MockHTTPClient is a mock implementation of HTTPClient for testing
type MockHTTPClient struct {
	MockResponse *http.Response
	MockError    error
}

// Get performs a mocked GET request
func (m *MockHTTPClient) Get(url string) (*http.Response, error) {
	return m.MockResponse, m.MockError
}

// Example usage
func main() {
	// Production usage
	realClient := &RealHTTPClient{}
	apiClient := NewAPIClient(realClient)

	user, err := apiClient.GetUser(1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("User: %+v\n", user)
}

*/
