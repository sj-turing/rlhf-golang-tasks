package main

/*

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// Simulates an unreliable service
func unreliableService() error {
	// Simulating a failure with a 70% chance
	if rand.Intn(100) < 70 {
		return fmt.Errorf("service unavailable")
	}
	return nil
}

// Retry Logic with Exponential Backoff
func callWithRetry(maxRetries int, delay time.Duration) error {
	var err error

	for attempts := 0; attempts < maxRetries; attempts++ {
		err = unreliableService()
		if err == nil {
			return nil // Success
		}

		// Exponential backoff
		time.Sleep(delay)
		delay *= 2 // Double the wait time for the next attempt
	}

	return err
}

// Circuit Breaker Implementation
type CircuitBreaker struct {
	failureThreshold int
	retryTimeout     time.Duration
	failureCount     int
	state            string // can be 'CLOSED', 'OPEN', 'HALF-OPEN'
	lastFailureTime  time.Time
}

func NewCircuitBreaker(threshold int, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		failureThreshold: threshold,
		retryTimeout:     timeout,
		state:            "CLOSED",
	}
}

func (cb *CircuitBreaker) callService() error {
	if cb.state == "OPEN" {
		if time.Since(cb.lastFailureTime) < cb.retryTimeout {
			return fmt.Errorf("circuit breaker open: service unavailable")
		}
		cb.state = "HALF-OPEN" // Try to call the service again
	}

	err := unreliableService()
	if err != nil {
		cb.failureCount++
		cb.lastFailureTime = time.Now()
		if cb.failureCount >= cb.failureThreshold {
			cb.state = "OPEN"
		}
	} else {
		cb.failureCount = 0 // Reset on success
	}

	return err
}

func main() {
	// Demonstrating Retry Logic with Exponential Backoff
	fmt.Println("Attempting to call the unreliable service with retry logic...")

	if err := callWithRetry(5, 1*time.Second); err != nil {
		fmt.Printf("Final error after retries: %s\n", err)
	} else {
		fmt.Println("Successfully called the service!")
	}

	fmt.Println("---------------------------------------------------")

	// Demonstrating Circuit Breaker
	cb := NewCircuitBreaker(3, 5*time.Second)

	fmt.Println("Attempting to call the service with a circuit breaker...")
	for i := 0; i < 10; i++ {
		if err := cb.callService(); err != nil {
			fmt.Printf("Error: %s\n", err)
		} else {
			fmt.Println("Successfully called the service!")
		}
		time.Sleep(1 * time.Second) // Adding a delay between calls
	}

}
*/
