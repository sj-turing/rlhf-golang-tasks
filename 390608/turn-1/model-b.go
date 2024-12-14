package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// exponentialBackoff calculates the backoff duration for retry attempts
func exponentialBackoff(attempt int, maxDelay time.Duration) time.Duration {
	base := time.Second * 1
	multiplier := float64(1 << attempt) // 2^attempt
	delay := time.Duration(min(float64(maxDelay), multiplier))
	// Add a jitter factor to prevent synchronization attacks
	jitter := time.Duration(rand.Float64() * float64(delay))
	return delay + jitter
}

// min returns the minimum of two values
func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

// retryHTTPClient sends an HTTP request with retries and exponential backoff
func retryHTTPClient(ctx context.Context, client *http.Client, url string, maxAttempts int, maxDelay time.Duration) (*http.Response, error) {
	for attempt := 0; attempt < maxAttempts; attempt++ {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create request: %w", err)
		}

		resp, err := client.Do(req)
		if err != nil {
			// Determine if we should retry
			if attempt < maxAttempts-1 {
				select {
				case <-ctx.Done():
					return nil, ctx.Err()
				default:
					backoff := exponentialBackoff(attempt, maxDelay)
					fmt.Printf("Attempt %d failed, retrying in %s...\n", attempt+1, backoff)
					time.Sleep(backoff)
				}
			}
			return nil, fmt.Errorf("failed to get response: %w", err)
		}

		// Check response status
		if resp.StatusCode >= 500 {
			// Retry on server errors
			resp.Body.Close()
			if attempt < maxAttempts-1 {
				select {
				case <-ctx.Done():
					return nil, ctx.Err()
				default:
					backoff := exponentialBackoff(attempt, maxDelay)
					fmt.Printf("Server error (%d), retrying in %s...\n", resp.StatusCode, backoff)
					time.Sleep(backoff)
				}
			}
			return nil, fmt.Errorf("server error: %d", resp.StatusCode)
		}

		defer resp.Body.Close()
		return resp, nil
	}

	return nil, fmt.Errorf("maximum retries reached")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ctx := context.Background()

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	url := "https://api.example.com/data" // Replace with your API endpoint
	maxAttempts := 5
	maxDelay := time.Minute * 5

	resp, err := retryHTTPClient(ctx, client, url, maxAttempts, maxDelay)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Status:", resp.Status)
}
