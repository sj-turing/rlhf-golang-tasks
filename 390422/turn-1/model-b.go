package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/pkg/errors"
)

// Simulate a function that might fail
func doWork() error {
	time.Sleep(1 * time.Second) // Simulate work
	// Introduce a random error
	if rand.Intn(2) == 0 {
		return errors.New("work failed randomly")
	}
	return nil
}

// A goroutine that performs work and returns errors via a channel
func worker(id int, errCh chan error) {
	defer close(errCh)
	log.Printf("Worker %d started\n", id)
	err := doWork()
	if err != nil {
		// Wrap the error to add context
		err = errors.Wrapf(err, "worker %d failed", id)
		errCh <- err
	} else {
		errCh <- nil
	}
	log.Printf("Worker %d finished\n", id)
}

func main() {
	// Create a channel to receive errors from goroutines
	errCh := make(chan error, 2)

	// Start two goroutines
	go worker(1, errCh)
	go worker(2, errCh)

	// Collect errors from both goroutines
	var errorsOccurred []error
	for i := 0; i < 2; i++ {
		if err := <-errCh; err != nil {
			errorsOccurred = append(errorsOccurred, err)
		}
	}

	// Check if any errors occurred
	if len(errorsOccurred) > 0 {
		for _, err := range errorsOccurred {
			fmt.Println(err) // Outputs the full error stack with context
		}
	} else {
		fmt.Println("All workers completed successfully.")
	}
}
