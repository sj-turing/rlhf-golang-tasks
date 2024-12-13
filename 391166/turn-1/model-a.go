package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

// ProcessData simulates the processing of a single data item.
func ProcessData(ctx context.Context, data int) error {
	// Simulate a random error
	if rand.Intn(10) < 3 { // 30% chance of error
		return fmt.Errorf("failed to process data %d", data)
	}
	return nil
}

// StreamData processes a stream of integers and handles errors.
func StreamData(ctx context.Context, dataStream []int) {
	var errors []error
	for _, data := range dataStream {
		select {
		case <-ctx.Done():
			log.Println("Stream processing was canceled.")
			return
		default:
			err := ProcessData(ctx, data)
			if err != nil {
				log.Printf("Error processing data %d: %v\n", data, err)
				errors = append(errors, err)
			}
		}
	}

	// Report all errors encountered
	if len(errors) > 0 {
		log.Printf("Encountered %d errors during processing:\n", len(errors))
		for _, err := range errors {
			log.Println(err)
		}
	} else {
		log.Println("All data processed successfully.")
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Create a cancelable context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Simulating a stream of data
	dataStream := make([]int, 100)
	for i := 0; i < 100; i++ {
		dataStream[i] = i
	}

	StreamData(ctx, dataStream)
}
