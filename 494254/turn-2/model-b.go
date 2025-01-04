package main

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// fileSize holds the name and size of a file
type fileSize struct {
	Name string
	Size int64
}

// walkDirFunc is a file walk function that calculates file sizes
func walkDirFunc(ctx context.Context, fileSizes chan<- fileSize, dirPath string, errChan chan<- error) filepath.WalkDirFunc {
	return func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		// Check for context cancellation
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		// Calculate the file size and send it to the channel
		info, err := os.Stat(path)
		if err != nil {
			errChan <- fmt.Errorf("error stat'ing file %q: %w", path, err)
			return nil
		}
		fileSizes <- fileSize{Name: path, Size: info.Size()}
		return nil
	}
}

// sumFileSizes calculates the total size of files from the fileSizes channel
func sumFileSizes(ctx context.Context, fileSizes <-chan fileSize, totalSize *int64, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case fs := <-fileSizes:
			(*totalSize) += fs.Size
		}
	}
}

func main() {
	// Command line argument handling
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./file_size_calc <path>")
		os.Exit(1)
	}
	path := os.Args[1]

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Channel to receive file sizes
	fileSizes := make(chan fileSize)
	// Channel to handle errors
	errChan := make(chan error, 1)

	// Number of goroutines for file size summation
	const numWorkers = 4
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Start workers to calculate file sizes concurrently
	for i := 0; i < numWorkers; i++ {
		go sumFileSizes(ctx, fileSizes, &totalSize, &wg)
	}

	// Walk the directory and send file sizes to the channel
	walkErr := filepath.WalkDir(path, walkDirFunc(ctx, fileSizes, path, errChan))
	close(fileSizes)

	// Wait for all workers to finish
	wg.Wait()

	// Select on the error channel to handle any errors
	select {
	case err := <-errChan:
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	case <-ctx.Done():
		fmt.Println("Calculation timed out. Please try again later.")
	default:
		if walkErr != nil {
			fmt.Printf("Error walking directory %q: %v\n", path, walkErr)
			os.Exit(1)
		}
		// Print the total size of the files
		fmt.Printf("Total size of files in %q: %d bytes\n", path, totalSize)
	}
}
