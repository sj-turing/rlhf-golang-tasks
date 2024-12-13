package main

/*
import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

const (
	// Define a reasonable timeout for the context.
	streamTimeout = 10 * time.Second
)

// Custom error types for detailed error handling.
type CustomError struct {
	Op   string    // Operation that failed
	Err  error     // Underlying error
	Time time.Time // Time of occurrence
}

// Implement the error interface for CustomError.
func (ce *CustomError) Error() string {
	return fmt.Sprintf("error during %s: %v at %v", ce.Op, ce.Err, ce.Time)
}

// Function to process a line of data.
func processLine(line []byte) {
	// Simulate line processing.
	if string(line) == "error" {
		// Simulate a processing error.
		log.Printf("simulated error processing line: %s", string(line))
	} else {
		log.Printf("processing line: %s", string(line))
		time.Sleep(100 * time.Millisecond) // Simulating processing time.
	}
}

// Function to stream data from a file.
func streamData(ctx context.Context, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return &CustomError{"open file", err, time.Now()}
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("error closing file: %v", err)
		}
	}()

	reader := bufio.NewReader(file)

	for {
		select {
		case <-ctx.Done():
			log.Println("streaming context cancelled")
			return ctx.Err()
		default:
			line, isPrefix, err := reader.ReadLine()
			if err != nil {
				if err == io.EOF {
					log.Println("end of file reached")
					return nil
				}
				return &CustomError{"read line", err, time.Now()}
			}

			if isPrefix {
				log.Printf("warning: incomplete line read: %s", line)
			}

			processLine(line)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), streamTimeout)
	defer cancel()

	filePath := "large_data.txt" // Assume this file exists.

	err := streamData(ctx, filePath)
	if err != nil {
		log.Printf("error streaming data: %v", err)
	}
}

*/
