package main

import (
	"log"
	"testing"
)

func BenchmarkCreateWithConcurrency(b *testing.B) {
	tempDir := "./bench-test"

	b.StartTimer()
	var count = int32(len(tempDir))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		dirChan := generateDirectoryName(tempDir, count)
		go func() {
			for err := range batchProcess("create", dirChan) {
				if err != nil {
					log.Println(err)
				}
			}
		}()
	}
	b.StopTimer()
}

func BenchmarkDeleteWithConcurrency(b *testing.B) {
	tempDir := "./bench-test"
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		dirChan := getDirectoryNames(tempDir)
		go func() {
			for err := range batchProcess("delete", dirChan) {
				if err != nil {
					log.Println(err)
				}
			}
		}()
	}

	b.StopTimer()
}
