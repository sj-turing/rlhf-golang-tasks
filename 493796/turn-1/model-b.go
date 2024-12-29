package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sync"
)

// File interface defines the basic file operations.
type File interface {
	io.Reader
	io.Writer
	Close() error
}

type BasicFile struct {
	path string
	f    *os.File
	mu   sync.Mutex
}

func NewBasicFile(path string) (*BasicFile, error) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return &BasicFile{path: path, f: f}, nil
}

func (b *BasicFile) Read(p []byte) (int, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.f.Read(p)
}

func (b *BasicFile) Write(p []byte) (int, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.f.Write(p)
}

func (b *BasicFile) Close() error {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.f.Close()
}

type VersionedFile struct {
	File
	history []string
}

func NewVersionedFile(path string) (*VersionedFile, error) {
	f, err := NewBasicFile(path)
	if err != nil {
		return nil, err
	}
	return &VersionedFile{File: f}, nil
}

func (v *VersionedFile) Write(p []byte) (int, error) {
	n, err := v.File.Write(p)
	if err != nil {
		return n, err
	}
	// After writing, add the current content to history
	content, err := ioutil.ReadFile(v.File.(*BasicFile).path)
	if err != nil {
		return n, err
	}
	v.history = append(v.history, string(content))
	return n, nil
}

func (v *VersionedFile) GetHistory() []string {
	return v.history
}

// Decorator function to wrap a File with VersionedFile.
func WithVersioning(f File) File {
	return &VersionedFile{File: f}
}

func main() {
	filePath := "test.txt"

	// Open a basic file for concurrent modifications
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
			vf.Write([]byte(fmt.Sprintf("Modification %d\n", i)))
		}
	}()

	go func() {
		for i := 5; i < 10; i++ {
			vf.Write([]byte(fmt.Sprintf("Modification %d\n", i)))
		}
	}()

	// Wait for modifications to complete
	// In a real-world scenario, you would use appropriate synchronization.
	time.Sleep(2 * time.Second)

	// Get the version history
	history := vf.(*VersionedFile).GetHistory()
	fmt.Println("Version History:")
	for i, content := range history {
		fmt.Printf("Version %d:\n%s\n", i+1, content)
	}
}
