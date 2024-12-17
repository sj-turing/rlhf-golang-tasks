package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Option represents a functional option for the NewAPI function.
type Option func(*API)

// API represents our concurrent API.
type API struct {
	url     string
	client  *http.Client
	workers int
	mutex   sync.Mutex
}

// WithURL sets the URL for API requests.
func WithURL(url string) Option {
	return func(a *API) {
		a.url = url
	}
}

// WithClient sets the HTTP client for API requests.
func WithClient(client *http.Client) Option {
	return func(a *API) {
		a.client = client
	}
}

// WithWorkers sets the number of worker Goroutines.
func WithWorkers(workers int) Option {
	return func(a *API) {
		a.workers = workers
	}
}

// NewAPI creates a new API instance with the given options.
func NewAPI(opts ...Option) *API {
	a := &API{
		url:     "https://jsonplaceholder.typicode.com/todos/1",
		client:  &http.Client{},
		workers: 1,
	}
	for _, opt := range opts {
		opt(a)
	}
	return a
}

// ProcessRequest sends a request to the API URL and processes the response.
func (a *API) ProcessRequest() {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	req, err := http.NewRequest("GET", a.url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	resp, err := a.client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Response status:", resp.Status)
}

// Serve initiates the API processing with Goroutines.
func (a *API) Serve() {
	for i := 0; i < a.workers; i++ {
		go func() {
			for {
				a.ProcessRequest()
				time.Sleep(time.Second)
			}
		}()
	}
}

func main() {
	// Create the API with custom options
	api := NewAPI(
		WithURL("https://jsonplaceholder.typicode.com/todos/10"),
		WithClient(&http.Client{Timeout: 5 * time.Second}),
		WithWorkers(3),
	)
	// Start the API processing
	api.Serve()
	// Block the main Goroutine to prevent exit
	select {}
}
