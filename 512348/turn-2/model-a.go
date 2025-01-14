package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

const dirSep = string(os.PathSeparator)

func createDir(dir string, wg *sync.WaitGroup, errors chan error) {
	defer wg.Done()
	defer close(errors)

	err := os.MkdirAll(dir, 0755)
	if err != nil {
		errors <- fmt.Errorf("error creating directory '%s': %w", dir, err)
	}
}

func deleteDir(dir string, wg *sync.WaitGroup, errors chan error) {
	defer wg.Done()
	defer close(errors)

	err := os.RemoveAll(dir)
	if err != nil {
		errors <- fmt.Errorf("error deleting directory '%s': %w", dir, err)
	}
}

func processBatches(mode string, batchSize int, dirs ...string) {
	for start := 0; start < len(dirs); start += batchSize {
		end := start + batchSize
		if end > len(dirs) {
			end = len(dirs)
		}
		wg := &sync.WaitGroup{}
		errors := make(chan error, end-start)

		for _, dir := range dirs[start:end] {
			wg.Add(1)

			if mode == "create" {
				go createDir(dir, wg, errors)
			} else {
				go deleteDir(dir, wg, errors)
			}
		}

		wg.Wait()
		close(errors)

		for err := range errors {
			log.Println(err)
		}
	}
}

func main() {
	if len(os.Args) < 4 || len(os.Args) > 5 {
		fmt.Println("Usage: go run main.go <mode> <root_dir> <batch_size> <directory_count>")
		fmt.Println("Example: go run main.go create ./bulk_dirs 10 100")
		fmt.Println("Example: go run main.go delete ./bulk_dirs 10")
		return
	}

	mode := os.Args[1]
	rootDir := os.Args[2]
	batchSizeStr := os.Args[3]

	batchSize, err := strconv.Atoi(batchSizeStr)
	if err != nil {
		log.Fatalf("Invalid batch size '%s': %v", batchSizeStr, err)
	}

	switch mode {
	case "create":
		directoryCountStr := os.Args[4]

		directoryCount, err := strconv.Atoi(directoryCountStr)
		if err != nil {
			log.Fatalf("Invalid directory size '%s': %v", directoryCountStr, err)
		}

		// Create a directory hierarchy
		dirs := make([]string, directoryCount)

		for index := range dirs {
			dirs[index] = rootDir + dirSep + fmt.Sprintf("%d", index)
		}

		processBatches(mode, batchSize, dirs...)
	case "delete":
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

		processBatches("delete", batchSize, dirs...)

	default:
		log.Fatalf("Invalid mode '%s'. Mode must be 'create' or 'delete'.", mode)
	}
}
