package main

import (
	"fmt"
	"os"
)

func main() {
	// Example usage with `os.Stat`
	fileName := "example.txt"

	// Stat the file to get its information and an error
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		// Check and handle the error
		if os.IsNotExist(err) {
			fmt.Printf("Error: File %s does not exist.\n", fileName)
		} else {
			fmt.Printf("An unexpected error occurred: %v\n", err)
		}
		return
	}

	// Print file information
	fmt.Printf("File mode: %v\n", fileInfo.Mode())
	fmt.Printf("File size: %d bytes\n", fileInfo.Size())
	fmt.Printf("File last modified: %s\n", fileInfo.ModTime())

	// Example usage with `os.ReadDir`
	dirName := "."

	// Read the directory contents
	files, err := os.ReadDir(dirName)
	if err != nil {
		// Check and handle the error
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}

	// Print directory contents
	fmt.Printf("Directory contents:\n")
	for _, file := range files {
		fmt.Printf("%s\n", file.Name())
	}
}
