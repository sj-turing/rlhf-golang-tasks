package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync/atomic"
)

// Platform-specific directory separator
const dirSep = string(os.PathSeparator)

// createDirs creates directories with the given batch
func createDirs(dir string, errChan chan<- error) {
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		errChan <- fmt.Errorf("Error creating directory '%s': %v", dir, err)
	}
	errChan <- nil
}

// deleteDirs deletes directories with the given batch
func deleteDirs(dir string, errChan chan<- error) {
	err := os.Remove(dir)
	if err != nil {
		errChan <- fmt.Errorf("Error deleting directory '%s': %v", dir, err)
	}
	errChan <- nil
}

// batchProcess batches the directory creation or deletion operations based on the provided mode.
func batchProcess(mode string, dirChan <-chan string) <-chan error {
	errChan := make(chan error)

	go func() {
		defer close(errChan)
		for {
			select {
			case dir, ok := <-dirChan:
				if !ok {
					return
				}
				if mode == "create" {
					createDirs(dir, errChan)
				} else {
					deleteDirs(dir, errChan)
				}
			case err, ok := <-errChan:
				if !ok {
					return
				}
				log.Println(err)
			}
		}
	}()

	return errChan
}

func generateDirectoryName(rootPath string, count int32) <-chan string {
	var bufferSize int32 = count
	if count > 100 {
		bufferSize = count / 10
	}
	directoryNameChan := make(chan string, bufferSize)

	go func() {
		defer close(directoryNameChan)
		for {
			if atomic.LoadInt32(&count) <= 0 {
				return
			}
			select {
			case directoryNameChan <- rootPath + dirSep + fmt.Sprintf("%d", count):
			}

			atomic.AddInt32(&count, -1)
		}
	}()

	return directoryNameChan
}

func getDirectoryNames(rootPath string) <-chan string {

	allFiles, err := os.ReadDir(rootPath)
	if err != nil {
		log.Fatalf("error while getting all files from the given path: %v\n", err)
	}
	var bufferSize int = len(allFiles)

	if len(allFiles) > 100 {
		bufferSize = len(allFiles) / 10
	}

	directoryNameChan := make(chan string, bufferSize)

	go func() {
		defer close(directoryNameChan)
		for _, file := range allFiles {
			directoryNameChan <- rootPath + dirSep + file.Name()
		}
	}()

	return directoryNameChan
}

// main function of the bulk directory management program.
func main() {
	if len(os.Args) < 3 || len(os.Args) > 4 {
		fmt.Println("Usage: go run main.go <mode> <root_dir> <directory_count>")
		fmt.Println("Example: go run main.go create ./bulk_dirs 100")
		fmt.Println("Example: go run main.go delete ./bulk_dirs")
		return
	}

	mode := os.Args[1]
	rootDir := os.Args[2]

	var directoryNameChan <-chan string

	switch mode {
	case "create":
		directoryCountStr := os.Args[3]

		directoryCount, err := strconv.Atoi(directoryCountStr)
		if err != nil {
			log.Fatalf("Invalid directory size '%s': %v", directoryCountStr, err)
		}

		directoryNameChan = generateDirectoryName(rootDir, int32(directoryCount))

	case "delete":
		directoryNameChan = getDirectoryNames(rootDir)

	default:
		log.Fatalf("Invalid mode '%s'. Mode must be 'create' or 'delete'.", mode)
	}

	for err := range batchProcess(mode, directoryNameChan) {
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("Program executed!")

}
