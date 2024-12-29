package main

/*

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

// File interface defines the basic file operations.
type File interface {
	io.Reader
	io.Writer
	Close() error
}

// BasicFile implements the File interface for basic file operations.
type BasicFile struct {
	*os.File
	mu sync.Mutex
}

func NewBasicFile(path string) (*BasicFile, error) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return &BasicFile{File: f}, nil
}

func (b *BasicFile) Read(p []byte) (int, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.File.Read(p)
}

func (b *BasicFile) Write(p []byte) (int, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.File.Write(p)
}

func (b *BasicFile) Close() error {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.File.Close()
}

// VersionedFile implements the File interface and decorates the BasicFile to track history.
type VersionedFile struct {
	File
	mu      sync.Mutex
	history []string
}

func NewVersionedFile(path string) (*VersionedFile, error) {
	f, err := NewBasicFile(path)
	if err != nil {
		return nil, err
	}
	return &VersionedFile{File: f, mu: sync.Mutex{}}, nil
}

func (v *VersionedFile) Write(p []byte) (int, error) {
	n, err := v.File.Write(p)
	if err != nil {
		return n, err
	}
	// After writing, add the current content to history
	content, err := ioutil.ReadFile(v.File.(*BasicFile).Name())
	if err != nil {
		return n, err
	}

	v.mu.Lock()
	v.history = append(v.history, string(content))
	v.mu.Unlock()
	return n, nil
}

func (v *VersionedFile) GetHistory() []string {
	v.mu.Lock()
	defer v.mu.Unlock()
	return v.history
}

// Decorator function to wrap a File with VersionedFile.
func WithVersioning(f File) File {
	return &VersionedFile{File: f}
}

func main() {
	filePath := "test.txt"

	// Open a basic file
	f, err := NewBasicFile(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	// Decorate the basic file with versioning
	vf := WithVersioning(f)

	// Perform concurrent modifications
	go func() {
		for i := 0; i < 5; i++ {
			_, err := vf.Write([]byte(fmt.Sprintf("Concurrent Modification %d\n", i)))
			if err != nil {
				fmt.Println("Error writing:", err)
			}
			time.Sleep(time.Duration(100*i) * time.Millisecond)
		}
	}()

	go func() {
		for i := 5; i < 10; i++ {
			_, err := vf.Write([]byte(fmt.Sprintf("Concurrent Modification %d\n", i)))
			if err != nil {
				fmt.Println("Error writing:", err)
			}
			time.Sleep(time.Duration(100*i) * time.Millisecond)
		}
	}()

	// Wait for modifications to complete
	time.Sleep(2 * time.Second)

	// Get the version history
	history := vf.(*VersionedFile).GetHistory()
	fmt.Println("Version History:")
	for i, content := range history {
		fmt.Printf("Version %d:\n%s\n", i+1, content)
	}
}

*/
