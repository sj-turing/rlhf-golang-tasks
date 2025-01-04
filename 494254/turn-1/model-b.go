package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func readFile(ctx context.Context, filename string) ([]byte, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func main() {
	// Create a context with a 5-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Read a file using the context
	data, err := readFile(ctx, "somefile.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(data))
}
