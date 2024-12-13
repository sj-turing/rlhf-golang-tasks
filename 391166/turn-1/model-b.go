package main

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
	// Set a reasonable timeout for the context
	streamTimeout = 5 * time.Second
)

func streamData(ctx context.Context, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
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
			log.Print("streaming context cancelled")
			return ctx.Err()
		default:
			line, isPrefix, err := reader.ReadLine()
			if err != nil {
				if err == io.EOF {
					log.Print("end of file reached")
					return nil
				}
				return fmt.Errorf("error reading line: %w", err)
			}

			if isPrefix {
				log.Printf("warning: incomplete line read: %s", line)
			}

			processLine(line)
		}
	}
}

func processLine(line []byte) {
	// Simulate processing logic
	log.Printf("processing line: %s", string(line))
	time.Sleep(100 * time.Millisecond) // Simulate processing time
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), streamTimeout)
	defer cancel()

	filePath := "large_data.txt"

	err := streamData(ctx, filePath)
	if err != nil {
		log.Printf("error streaming data: %v", err)
	}
}
