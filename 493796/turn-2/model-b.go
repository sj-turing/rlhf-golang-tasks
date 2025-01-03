package main

import (
	"fmt"
	"io"
	"io/ioutil"
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

// BasicFile implements the File interface using the os package.
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

// VersionedFile is a decorator that tracks historical changes of a BasicFile.
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
