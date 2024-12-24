package main

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

// Define a custom cleanup protocol
type Cleanup interface {
	Close() error
}

// Example of an external resource that needs cleanup
type FileCloser struct {
	*os.File
}

func (fc *FileCloser) Close() error {
	return fc.File.Close()
}

type DatabaseConnection struct {
	// Simulate database connection state
	isConnected bool
}

func (dc *DatabaseConnection) Close() error {
	dc.isConnected = false
	return nil
}

func (dc *DatabaseConnection) IsConnected() bool {
	return dc.isConnected
}

func main() {
	// Open a file
	file, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	fc := &FileCloser{file}
	defer fc.Close()

	// Connect to a database
	db := &DatabaseConnection{isConnected: true}
	defer db.Close()

	// Use the resources
	_, err = io.Copy(os.Stdout, fc)
	if err != nil {
		panic(err)
	}

	// Simulate some work using the database
	fmt.Println("Working with the database...")
	time.Sleep(2 * time.Second)

	if !db.IsConnected() {
		fmt.Println("Database connection is closed!")
	}

	// Acquire a lock and defer its release
	var wg sync.WaitGroup
	wg.Add(1)
	defer wg.Done()

	lock := &sync.Mutex{}
	lock.Lock()
	defer lock.Unlock()

	// Do some work with the lock held
	fmt.Println("Working with the lock...")

	// Wait for other goroutines to finish
	wg.Wait()
}
