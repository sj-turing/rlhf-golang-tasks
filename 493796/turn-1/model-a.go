package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

type FileSystem interface {
	ReadFile(string) ([]byte, error)
	WriteFile(string, []byte) error
}

type BasicFileSystem struct{}

func (bfs *BasicFileSystem) ReadFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func (bfs *BasicFileSystem) WriteFile(filename string, data []byte) error {
	return os.WriteFile(filename, data, 0644)
}

type HistoryEntry struct {
	Operation string // e.g., "write", "read"
	Filename  string
	Time      time.Time
	Content   []byte
}

type HistoryFileSystem struct {
	wrapped     FileSystem
	history     []HistoryEntry
	historyFile *os.File
}

func NewHistoryFileSystem(wrapped FileSystem, historyFilename string) (*HistoryFileSystem, error) {
	file, err := os.OpenFile(historyFilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	return &HistoryFileSystem{
		wrapped:     wrapped,
		history:     make([]HistoryEntry, 0),
		historyFile: file,
	}, nil
}

func (hfs *HistoryFileSystem) ReadFile(filename string) ([]byte, error) {
	data, err := hfs.wrapped.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	entry := HistoryEntry{
		Operation: "read",
		Filename:  filename,
		Time:      time.Now(),
		Content:   data,
	}

	hfs.history = append(hfs.history, entry)
	return data, nil
}

func (hfs *HistoryFileSystem) WriteFile(filename string, data []byte) error {
	err := hfs.wrapped.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	entry := HistoryEntry{
		Operation: "write",
		Filename:  filename,
		Time:      time.Now(),
		Content:   data,
	}

	hfs.history = append(hfs.history, entry)

	// Log entry to a history file
	_, err = hfs.historyFile.Write([]byte(fmt.Sprintf("%+v\n", entry)))
	if err != nil {
		return errors.Wrap(err, "failed to log history entry")
	}

	return nil
}

func main() {
	// Create basic file system instance
	basicFS := &BasicFileSystem{}

	// Create history file system wrapper
	historyFileName := "history.log"
	historyFS, err := NewHistoryFileSystem(basicFS, historyFileName)
	if err != nil {
		panic(err)
	}

	defer historyFS.historyFile.Close()

	// Modify the file
	fileName := "example.txt"
	data := []byte("Hello, world!")
	err = historyFS.WriteFile(fileName, data)
	if err != nil {
		panic(err)
	}

	// Read the file
	content, err := historyFS.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}
