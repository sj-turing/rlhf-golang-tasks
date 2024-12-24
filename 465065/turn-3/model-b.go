package main

import (
	"fmt"
	"io"
	"os"
)

// readFile reads a file and returns its contents as a string and an error.
func readFile(filename string) (string, error) {
	// Open the file for reading.
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Read the file contents.
	b, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func main() {
	// Call the readFile function and handle the error.
	content, err := readFile("example.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Print the file contents.
	fmt.Println(content)
}
