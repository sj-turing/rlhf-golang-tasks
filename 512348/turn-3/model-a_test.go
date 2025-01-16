package main

import (
	"testing"
)

func BenchmarkCreateNoConcurrency(b *testing.B) {
	b.StopTimer()
	dirSep := string(os.PathSeparator)
	rootDir := "./benchmark-dirs"
	fileCount := 10000

	os.MkdirAll(rootDir, 0755)
	defer os.RemoveAll(rootDir)

	dirs := make([]string, fileCount)
	for index := range dirs {
		dirs[index] = rootDir + dirSep + fmt.Sprintf("%d", index)
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		for _, dir := range dirs {
			os.MkdirAll(dir, 0755)
		}
	}

	b.StopTimer()
}

func BenchmarkDeleteNoConcurrency(b *testing.B) {
	b.StopTimer()
	dirSep := string(os.PathSeparator)
	rootDir := "./benchmark-dirs"
	fileCount := 10000

	os.MkdirAll(rootDir, 0755)
	defer os.RemoveAll(rootDir)

	dirs := make([]string, fileCount)
	for index := range dirs {
		dirs[index] = rootDir + dirSep + fmt.Sprintf("%d", index)
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		for _, dir := range dirs {
			os.Remove(dir)
		}
	}

	b.StopTimer()
}

func BenchmarkCreateWithConcurrency(b *testing.B) {
	b.StopTimer()
	dirSep := string(os.PathSeparator)
	rootDir := "./benchmark-dirs"
	fileCount := 10000

	os.MkdirAll(rootDir, 0755)
	defer os.RemoveAll(rootDir)

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		var count atomic.Int32
		dirChan := generateDirectoryName(rootDir, count.Load())
		c := make(chan error)
		go func() {
			for err := range batchProcess("create", dirChan) {
				if err != nil {
					log.Println(err)
				}
			}
		}()
		<-c
	}

	b.StopTimer()
}

func BenchmarkDeleteWithConcurrency(b *testing.B) {
	b.StopTimer()
	dirSep := string(os.PathSeparator)
	rootDir := "./benchmark-dirs"
	fileCount := 10000

	os.MkdirAll(rootDir, 0755)
	defer os.RemoveAll(rootDir)

	dirs := make([]string, fileCount)
	for index := range dirs {
		dirs[index] = rootDir + dirSep + fmt.Sprintf("%d", index)
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		var count atomic.Int32
		dirChan := generateDirectoryName(rootDir, count.Load())
		c := make(chan error)
		go func() {
			for err := range batchProcess("delete", dirChan) {
				if err != nil {
					log.Println(err)
				}
			}
		}()
		<-c
	}

	b.StopTimer()
}
