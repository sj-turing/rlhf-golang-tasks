package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

// Cleanup is a protocol for resource cleanup.
type Cleanup interface {
	// Close should be called to release resources.
	Close() error
}

// FileCleanup implements the Cleanup protocol for file resources.
type FileCleanup struct {
	*os.File
}

// Close closes the file and removes it.
func (fc *FileCleanup) Close() error {
	if err := fc.File.Close(); err != nil {
		return err
	}
	return os.Remove(fc.Name())
}

// NetworkCleanup implements the Cleanup protocol for network resources.
type NetworkCleanup struct {
	*os.File
}

// Close closes the network connection.
func (nc *NetworkCleanup) Close() error {
	return nc.File.Close()
}

func main() {
	var cleanupError error
	var wg sync.WaitGroup

	// Example: create a temporary file
	f, err := os.CreateTemp("", "example")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			cleanupError = fmt.Errorf("error closing file: %w", err)
		}
	}()

	fc := &FileCleanup{f}

	// Write something to the file
	fmt.Fprintf(f, "Hello, this is a temporary file!\n")

	// Defer the custom cleanup function
	defer func() {
		if err := fc.Close(); err != nil {
			cleanupError = fmt.Errorf("error closing file cleanup: %w", err)
		}
	}()

	fmt.Println("File contents:", string(mustReadAll(fc)))

	// Simulate an asynchronous operation
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Asynchronous operation complete...")
	}()

	// Wait for the asynchronous operation to complete
	wg.Wait()

	// If there was a cleanup error, report it
	if cleanupError != nil {
		fmt.Println("Cleanup error:", cleanupError)
	}
}

func mustReadAll(f Cleanup) []byte {
	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(fmt.Errorf("error reading from file: %w", err))
	}
	return data
}
