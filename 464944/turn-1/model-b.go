package main

import (
	"fmt"
	"io"
	"os"
)

// Define a custom cleanup protocol
type Cleanup interface {
	Close() error
}

// Example of an external resource that needs cleanup
type FileCloser struct {
	*os.File
}

// Implement the Close() method to satisfy the Cleanup protocol
func (fc *FileCloser) Close() error {
	return fc.File.Close()
}

func main() {
	// Open a file as an example of an external resource
	file, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	// Create a FileCloser to wrap the os.File and satisfy the Cleanup protocol
	fc := &FileCloser{file}

	// Use defer to ensure the file is closed automatically after the main function returns
	defer func() {
		if err := fc.Close(); err != nil {
			fmt.Println("Error closing file:", err)
		}
	}()

	// Now you can use the file as needed
	_, err = io.Copy(os.Stdout, fc)
	if err != nil {
		panic(err)
	}
}
