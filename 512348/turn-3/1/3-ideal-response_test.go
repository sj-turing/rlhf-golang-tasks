package main

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func BenchmarkCreateNoConcurrency(b *testing.B) {
	b.StopTimer()
	dirSep := string(os.PathSeparator)
	rootDir := "./benchmark-dirs"
	directoryCount := 100
	batchSize := 10

	dirs := make([]string, directoryCount)

	for index := range dirs {
		dirs[index] = rootDir + dirSep + fmt.Sprintf("%d", index)
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		batchProcess("create", batchSize, dirs...)
	}
	b.StopTimer()

}

func BenchmarkDeleteNoConcurrency(b *testing.B) {
	b.StopTimer()
	dirSep := string(os.PathSeparator)
	rootDir := "./benchmark-dirs"
	batchSize := 10

	allFiles, err := os.ReadDir(rootDir)
	if err != nil {
		log.Fatalf("error while getting all files from the given path: %v\n", err)
	}

	dirs := []string{}
	for _, file := range allFiles {
		if file.IsDir() {
			dirs = append(dirs, rootDir+dirSep+file.Name())
		}
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		batchProcess("delete", batchSize, dirs...)
	}

	b.StopTimer()
}
