package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Logger interface for different logging implementations
type Logger interface {
	Log(message string)
}

// ConsoleLogger logs to the console
type ConsoleLogger struct{}

func (c *ConsoleLogger) Log(message string) {
	fmt.Println("Console:", message)
}

// FileLogger logs to a file
type FileLogger struct {
	file *os.File
}

func (f *FileLogger) Log(message string) {
	fmt.Fprintln(f.file, "File:", message)
}

func NewFileLogger(filename string) (*FileLogger, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return &FileLogger{file: file}, nil
}

// GetLogger selects a logger based on the environment
func GetLogger() (Logger, error) {
	logType := os.Getenv("LOG_TYPE")
	switch logType {
	case "file":
		filename := os.Getenv("LOG_FILE")
		if filename == "" {
			return nil, fmt.Errorf("LOG_FILE environment variable not set")
		}
		return NewFileLogger(filename)
	default: // fallback to console logger
		return &ConsoleLogger{}, nil
	}
}

// HTTPClient interface for making HTTP requests
type HTTPClient interface {
	Get(url string) (*http.Response, error)
}

// RealHTTPClient makes actual HTTP requests
type RealHTTPClient struct{}

func (c *RealHTTPClient) Get(url string) (*http.Response, error) {
	return http.Get(url)
}

// APIClient consumes an HTTPClient and a Logger
type APIClient struct {
	httpClient HTTPClient
	logger     Logger
}

// NewAPIClient creates a new APIClient with the provided HTTPClient and Logger
func NewAPIClient(httpClient HTTPClient, logger Logger) *APIClient {
	return &APIClient{
		httpClient: httpClient,
		logger:     logger,
	}
}

// GetUser fetches a user by ID from the JSON API
func (c *APIClient) GetUser(id int) (*User, error) {
	resp, err := c.httpClient.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%d", id))
	if err != nil {
		c.logger.Log("Error fetching user: " + err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.logger.Log("Error fetching user: " + resp.Status)
		return nil, fmt.Errorf("error fetching user: %s", resp.Status)
	}

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		c.logger.Log("Error decoding user: " + err.Error())
		return nil, err
	}
	return &user, nil
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Example usage
func main() {
	logger, err := GetLogger()
	if err != nil {
		fmt.Println("Error initialising logger:", err)
		return
	}

	client := NewAPIClient(&RealHTTPClient{}, logger)

	user, err := client.GetUser(1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	logger.Log(fmt.Sprintf("%+v", user))
}
