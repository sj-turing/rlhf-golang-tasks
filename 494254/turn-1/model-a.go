package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"
)

func readFileWithContext(ctx context.Context, filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buffer := make([]byte, 1024) // Read in chunks of 1024 bytes
	var data []byte

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			n, err := file.Read(buffer)
			if n == 0 {
				return data, err
			}
			if err != nil && err != io.EOF {
				return nil, err
			}
			data = append(data, buffer[:n]...)
		}
	}
}

func main() {
	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filename := "example.txt"

	// Read the file using the context
	data, err := readFileWithContext(ctx, filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print the file content
	fmt.Println("File content:", string(data))
}
