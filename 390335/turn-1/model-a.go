package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"
)

// Custom error type to encapsulate error information
type StreamError struct {
	Code    int
	Message string
	Context string
}

func (e *StreamError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s, Context: %s", e.Code, e.Message, e.Context)
}

// DataStream represents a simple data source
type DataStream struct {
	data chan int
}

// NewDataStream creates a new data stream
func NewDataStream() *DataStream {
	return &DataStream{
		data: make(chan int),
	}
}

// Start simulates data production
func (ds *DataStream) Start() {
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			ds.data <- i
		}
		close(ds.data)
	}()
}

// Process simulates data processing and error handling
func Process(ctx context.Context, data <-chan int, errorChan chan<- error) {
	for val := range data {
		if err := doWork(val); err != nil {
			errorChan <- &StreamError{Code: 500, Message: err.Error(), Context: fmt.Sprintf("Processing value: %d", val)}
		}
	}
}

// doWork simulates processing work which can fail randomly
func doWork(value int) error {
	if rand.Float32() < 0.2 { // 20% chance to fail
		return fmt.Errorf("failed to process value: %d", value)
	}
	return nil
}

// ErrorReporter handles and logs errors
func ErrorReporter(errorChan <-chan error) {
	for err := range errorChan {
		// Here we would normally send alerts or log to a monitoring tool
		log.Printf("Error occurred: %s\n", err.Error())
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	ctx := context.Background()
	dataStream := NewDataStream()
	dataStream.Start()

	errorChan := make(chan error)

	go Process(ctx, dataStream.data, errorChan)
	go ErrorReporter(errorChan)

	// Wait for user input to exit
	fmt.Println("Press Enter to exit...")
	_, _ = fmt.Scanln()

	// Clean up
	close(errorChan)
}
