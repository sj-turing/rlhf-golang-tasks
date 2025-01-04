package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func calculateFileSize(path string, errChan chan<- error, fileSizeChan chan<- int64) {
	err := filepath.Walk(path, func(subpath string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error processing file %s: %v", subpath, err)
		}

		if info.IsDir() {
			return nil // Continue walking
		}

		if info.Mode()&os.ModeSymlink == os.ModeSymlink {
			return nil // Ignore symlinks
		}

		// Calculate the file size
		fileSizeChan <- info.Size()
		return nil
	})

	if err != nil {
		errChan <- err
	}
	close(fileSizeChan)
	close(errChan)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <directory-path>")
		os.Exit(1)
	}

	path := os.Args[1]

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	fileSizeChan := make(chan int64)
	errorChan := make(chan error, 1)

	go calculateFileSize(path, errorChan, fileSizeChan)

	var totalSize int64 = 0

forloop:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Calculation timed out. Please try again later.")
			os.Exit(1)
		case err, ok := <-errorChan:
			if ok {
				fmt.Println("Error while processing: ", err)
			}
		case size, ok := <-fileSizeChan:
			if !ok {
				break forloop
			}
			totalSize += size
		}
	}

	fmt.Printf("Total size of files in %s: %d bytes\n", path, totalSize)
}
