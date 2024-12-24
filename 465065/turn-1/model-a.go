package main

import (
	"fmt"
	"os"
)

func main() {
	// Define a file name
	fileName := "example.txt"

	// Create a file with specific permissions
	file, err := os.Create(fileName)
	if err != nil {
		// Check and handle the error
		fmt.Printf("Error creating file: %v\n", err)
		return
	}

	defer file.Close()

	// Write some data to the file
	n, err := file.Write([]byte("Hello, World!"))
	if err != nil {
		// Check and handle the error
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}

	// Check how many bytes were written
	fmt.Printf("%d bytes written to file.\n", n)

	// Stat the file to get its information
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		// Check and handle the error
		fmt.Printf("Error statting file: %v\n", err)
		return
	}

	// Print some file information
	fmt.Printf("File mode: %v\n", fileInfo.Mode())
	fmt.Printf("File size: %d bytes\n", fileInfo.Size())
	fmt.Printf("File last modified: %s\n", fileInfo.ModTime())
}
