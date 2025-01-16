package main

import (
	"os"
	"path/filepath"
	"strconv"
	"testing"
)

// BenchmarkBulkFileOpsWithoutConcurrency benchmarks the performance of bulk file operations without concurrency.
func BenchmarkBulkFileOpsWithoutConcurrency(b *testing.B) {
	tempDir := b.TempDir()
	defer os.RemoveAll(tempDir)

	for n := 0; n < b.N; n++ {
		for i := 0; i < 1000; i++ {
			filePath := filepath.Join(tempDir, strconv.Itoa(i))
			if err := os.WriteFile(filePath, []byte("test data"), 0644); err != nil {
				b.Fatalf("Error creating file: %v", err)
			}
			if err := os.Remove(filePath); err != nil {
				b.Fatalf("Error removing file: %v", err)
			}
		}
	}
}

// BenchmarkBulkFileOpsWithConcurrency benchmarks the performance of bulk file operations with concurrency.
func BenchmarkBulkFileOpsWithConcurrency(b *testing.B) {
	tempDir := b.TempDir()
	defer os.RemoveAll(tempDir)

	filePaths := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		filePaths[i] = filepath.Join(tempDir, strconv.Itoa(i))
	}

	for n := 0; n < b.N; n++ {
		ch := make(chan error, 1000)

		for _, filePath := range filePaths {
			go func(filePath string) {
				if err := os.WriteFile(filePath, []byte("test data"), 0644); err != nil {
					ch <- err
					return
				}
				if err := os.Remove(filePath); err != nil {
					ch <- err
					return
				}
				ch <- nil
			}(filePath)
		}

		for i := 0; i < 1000; i++ {
			if err := <-ch; err != nil {
				b.Fatalf("Error occurred during bulk file ops: %v", err)
			}
		}
	}
}
