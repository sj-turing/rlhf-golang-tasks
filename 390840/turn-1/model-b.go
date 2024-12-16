package main

import (
	"fmt"
	"io/ioutil"
	"testing"
	"time"
)

// DataStore is a functional interface with a single method.
type DataStore interface {
	LoadData(filename string) ([]byte, error)
}

// realDataStore is a concrete implementation of DataStore using ioutil.ReadFile.
type realDataStore struct{}

func (ds *realDataStore) LoadData(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

// mockDataStore is a mock implementation of DataStore for testing purposes.
type mockDataStore struct {
	data []byte
	err  error
}

func (ds *mockDataStore) LoadData(filename string) ([]byte, error) {
	return ds.data, ds.err
}

// ProcessData processes some data loaded from a DataStore.
func ProcessData(ds DataStore, filename string) error {
	data, err := ds.LoadData(filename)
	if err != nil {
		return err
	}
	fmt.Printf("Processed %d bytes\n", len(data))
	return nil
}

func BenchmarkProcessDataReal(b *testing.B) {
	ds := &realDataStore{}
	for n := 0; n < b.N; n++ {
		ProcessData(ds, "example.txt")
	}
}

func BenchmarkProcessDataMock(b *testing.B) {
	// Use a mock data store with some test data
	data := make([]byte, 1024*1024) // 1MB of data
	ds := &mockDataStore{data: data, err: nil}
	for n := 0; n < b.N; n++ {
		ProcessData(ds, "example.txt")
	}
}

func TestProcessDataReal(t *testing.T) {
	if err := ProcessData(&realDataStore{}, "example.txt"); err != nil {
		t.Errorf("ProcessData failed: %v", err)
	}
}

func TestProcessDataMock(t *testing.T) {
	// Test with a mock data store that simulates an error
	ds := &mockDataStore{data: nil, err: fmt.Errorf("mock error")}
	if err := ProcessData(ds, "example.txt"); err == nil {
		t.Error("ProcessData expected an error but got none")
	}
}

func main() {
	testing.Main()
}
