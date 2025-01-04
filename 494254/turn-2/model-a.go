package main

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

func calculateFileSize(ctx context.Context, path string) (int64, error) {
	size := int64(0)
	err := filepath.Walk(path, func(subpath string, info os.FileInfo, err error) error {
		if err != nil {
			if ctx.Err() != nil {
				return ctx.Err() // Propagate cancellation
			}
			return fmt.Errorf("error processing file %s: %v", subpath, err)
		}

		if info.IsDir() {
			return nil // Continue walking
		}

		if info.Mode()&os.ModeSymlink == os.ModeSymlink {
			return nil // Ignore symlinks
		}

		// Calculate the file size
		size += info.Size()
		return nil
	})
	return size, err
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <directory-path>")
		os.Exit(1)
	}

	path := os.Args[1]

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	totalSize, err := calculateFileSize(ctx, path)
	if err != nil {
		log.Fatalf("Error calculating file size: %v", err)
	}

	fmt.Printf("Total size of files in %s: %d bytes\n", path, totalSize)
}
