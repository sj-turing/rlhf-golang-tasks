package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// CountFilesInDirectory counts the number of files in a given directory
// and returns the count along with an error if any occurs.
func CountFilesInDirectory(path string) (int, error) {
	count := 0

	// Walk the directory tree
	err := filepath.Walk(path, func(entry string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		count++
		return nil
	})

	return count, err
}

func main() {
	// Define a directory path
	dirPath := "./"

	// Call the function to count files
	fileCount, err := CountFilesInDirectory(dirPath)
	if err != nil {
		// Check and handle the error
		fmt.Printf("Error counting files: %v\n", err)
		return
	}

	// Print the number of files found
	fmt.Printf("Number of files in directory: %d\n", fileCount)
}
