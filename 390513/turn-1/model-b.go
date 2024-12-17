package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Option represents a functional option for the fetchData function.
type Option func(f *Fetcher)

// Fetcher encapsulates the options for fetching data.
type Fetcher struct {
	timeout time.Duration
	retries int
}

// WithTimeout sets the timeout for the fetch operation.
func WithTimeout(t time.Duration) Option {
	return func(f *Fetcher) {
		f.timeout = t
	}
}

// WithRetries sets the number of retries for the fetch operation.
func WithRetries(r int) Option {
	return func(f *Fetcher) {
		f.retries = r
	}
}

// fetchData fetches data from a remote source with the given options.
func fetchData(url string, opts ...Option) ([]byte, error) {
	f := &Fetcher{
		timeout: 10 * time.Second, // Default timeout
		retries: 3,                // Default retries
	}

	// Apply functional options
	for _, opt := range opts {
		opt(f)
	}

	// Perform the fetch operation
	return fetchDataWithOptions(url, f)
}

func fetchDataWithOptions(url string, f *Fetcher) ([]byte, error) {
	// Customize HTTP client with options
	client := &http.Client{
		Timeout: f.timeout,
	}

	for i := 0; i < f.retries+1; i++ {
		resp, err := client.Get(url)
		if err != nil {
			// Handle network errors
			continue
		}

		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			return ioutil.ReadAll(resp.Body)
		}

		// Handle non-200 status codes
	}

	return nil, fmt.Errorf("fetch failed after %d retries", f.retries)
}

func main() {
	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		// Fetch data with custom options (e.g., timeout and retries)
		data, err := fetchData("https://jsonplaceholder.typicode.com/todos/1", WithTimeout(5*time.Second), WithRetries(2))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(data)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
