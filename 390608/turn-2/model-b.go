// filename: main.go
package main

/*
import (
	"context"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

func main() {
	ctx := context.Background()

	// Example of retry logic with exponential backoff
	log.Println("Running retry logic with exponential backoff...")
	if err := makeHTTPRequestWithRetry(ctx, "http://localhost:9000/retry-example"); err != nil {
		log.Fatalf("Failed after retries: %v", err)
	}
	log.Println("Retry logic succeeded.")

	// Example of circuit breaker
	log.Println("Running circuit breaker...")
	circuitBreaker := newCircuitBreaker(3, 5*time.Second, 2)
	if err := makeHTTPRequestWithCircuitBreaker(ctx, circuitBreaker, "http://localhost:9000/circuit-breaker-example"); err != nil {
		log.Fatalf("Circuit breaker failed: %v", err)
	}
	log.Println("Circuit breaker succeeded.")
}

// Simulate a random failure for demonstration purposes
func simulateFailure() bool {
	return rand.Intn(5) == 0 // 20% chance of failure
}

func makeHTTPRequestWithRetry(ctx context.Context, url string) error {
	maxRetries := 5
	delay := 1 * time.Second

	for retry := 0; retry < maxRetries; retry++ {
		if err := makeHTTPRequest(ctx, url); err == nil {
			return nil // Success
		}

		log.Printf("Retry %d failed: %v. Retrying in %v...\n", retry+1, err, delay)
		time.Sleep(delay)
		delay *= 2
	}
	return fmt.Errorf("max retries reached")
}

func makeHTTPRequest(ctx context.Context, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if simulateFailure() {
		return fmt.Errorf("simulated failure")
	}

	return nil
}

type circuitBreaker struct {
	maxFailures     int
	resetTimeout    time.Duration
	successes       int
	failures        int
	lastFailureTime time.Time
	mu              sync.Mutex
	isOpen          bool
}

func newCircuitBreaker(maxFailures int, resetTimeout time.Duration, successes int) *circuitBreaker {
	return &circuitBreaker{
		maxFailures:     maxFailures,
		resetTimeout:    resetTimeout,
		successes:       successes,
		failures:        0,
		lastFailureTime: time.Now(),
		mu:              sync.Mutex{},
		isOpen:          false,
	}
}

func (cb *circuitBreaker) IsOpen() bool {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	return cb.isOpen
}

func (cb *circuitBreaker) markSuccess() {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	cb.successes++
	cb.failures = 0
	cb.lastFailureTime = time.Now()
	cb.isOpen = false
}

func (cb *circuitBreaker) markFailure() {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	cb.failures++
	cb.successes = 0
	cb.lastFailureTime = time.Now()

	if cb.failures >= cb.maxFailures {
		cb.isOpen = true
		go func() {
			time.Sleep(cb.resetTimeout)
			cb.markSuccess()
		}()
	}
}

func makeHTTPRequestWithCircuitBreaker(ctx context.Context, cb *circuitBreaker, url string) error {
	if cb.IsOpen() {
		return fmt.Errorf("circuit breaker is open")
	}

	if err := makeHTTPRequest(ctx, url); err != nil {
		cb.markFailure()
		return err
	}

	cb.markSuccess()
	return nil
}

*/
