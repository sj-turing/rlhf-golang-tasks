package main

/*

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

// Custom error type to encapsulate detailed error information
type StreamError struct {
	Code    int
	Message string
	Context string
}

func (e *StreamError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s, Context: %s", e.Code, e.Message, e.Context)
}

// DataStream represents a simple data source.
type DataStream struct {
	data chan int
}

// NewDataStream creates a new instance of DataStream.
func NewDataStream(size int) *DataStream {
	return &DataStream{
		data: make(chan int, size),
	}
}

// Start simulates data production.
func (ds *DataStream) Start() {
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			ds.data <- i
			log.Printf("Produced: %d\n", i)
		}
		close(ds.data)
	}()
}

// Process simulates data processing and error handling.
func Process(ctx context.Context, data <-chan int, errorChan chan<- error) {
	for val := range data {
		if err := doWork(val); err != nil {
			// Send detailed error information to the error channel
			errorChan <- &StreamError{Code: 500, Message: err.Error(), Context: fmt.Sprintf("Processing value: %d", val)}
		} else {
			log.Printf("Successfully processed: %d\n", val)
		}
	}
}

// doWork simulates a task which can fail.
func doWork(value int) error {
	// Randomly cause a failure
	if rand.Float32() < 0.3 { // 30% chance to fail
		return fmt.Errorf("failed to process value: %d", value)
	}
	return nil
}

// ErrorReporter handles and logs errors to a designated logging system.
func ErrorReporter(errorChan <-chan error) {
	for err := range errorChan {
		// Here you could log to a file or monitoring system
		log.Printf("Error occurred: %s\n", err.Error())
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Create a context for cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Set up data stream
	dataStream := NewDataStream(5)
	dataStream.Start()

	// Channel for reporting errors
	errorChan := make(chan error)

	// Start the error reporter.
	go ErrorReporter(errorChan)

	// Process the data
	Process(ctx, dataStream.data, errorChan)

	// Close the error channel after processing is complete
	close(errorChan)

	// Wait for user input before exiting
	fmt.Println("Press Enter to exit...")
	_, _ = fmt.Scanln()

	log.Println("Exiting application.")
	os.Exit(0)
}

*/
