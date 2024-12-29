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
}

func NewBasicFile(path string) (*BasicFile, error) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return &BasicFile{path: path, f: f}, nil
}

func (b *BasicFile) Read(p []byte) (int, error) {
	return b.f.Read(p)
}

func (b *BasicFile) Write(p []byte) (int, error) {
	return b.f.Write(p)
}

func (b *BasicFile) Close() error {
	return b.f.Close()
}

type VersionedFile struct {
	File
	history []string
	mu      *sync.Mutex
}

func NewVersionedFile(file File) (*VersionedFile, error) {
	return &VersionedFile{File: file, history: []string{}, mu: &sync.Mutex{}}, nil
}

func (v *VersionedFile) Write(p []byte) (int, error) {
	n, err := v.File.Write(p)
	if err != nil {
		return n, err
	}

	// Read the file content and append to history
	content, err := ioutil.ReadFile(v.File.(*BasicFile).path)
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

func main() {
	filePath := "test.txt"

	// Open a basic file
	basicFile, err := NewBasicFile(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer basicFile.Close()

	// Decorate the basic file with versioning
	versionedFile, err := NewVersionedFile(basicFile)
	if err != nil {
		fmt.Println("Error decorating file:", err)
		return
	}

	// Perform concurrent modifications
	go func() {
		for i := 0; i < 5; i++ {
			versionedFile.Write([]byte(fmt.Sprintf("Modification %d\n", i)))
		}
	}()

	go func() {
		for i := 5; i < 10; i++ {
			versionedFile.Write([]byte(fmt.Sprintf("Modification %d\n", i)))
		}
	}()

	// Wait for modifications to complete
	time.Sleep(2 * time.Second)

	// Get the version history
	history := versionedFile.GetHistory()
	fmt.Println("Version History:")
	for i, content := range history {
		fmt.Printf("Version %d:\n%s\n", i+1, content)
	}
}
