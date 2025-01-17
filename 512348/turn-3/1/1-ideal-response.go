package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// Platform-specific directory separator
const dirSep = string(os.PathSeparator)

// createDirs creates directories with the given batch
func createDirs(dirs []string) {
	for _, dir := range dirs {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			log.Printf("Error creating directory '%s': %v", dir, err)
		}
	}
}

// deleteDirs deletes directories with the given batch
func deleteDirs(dirs []string) {
	for _, dir := range dirs {
		err := os.RemoveAll(dir)
		if err != nil {
			log.Printf("Error deleting directory '%s': %v", dir, err)
		}
	}
}

// batchProcess batches the directory creation or deletion operations based on the provided mode.
func batchProcess(mode string, batchSize int, dirs ...string) {
	for start := 0; start < len(dirs); start += batchSize {
		end := start + batchSize
		if end > len(dirs) {
			end = len(dirs)
		}
		if mode == "create" {
			createDirs(dirs[start:end])
		} else {
			deleteDirs(dirs[start:end])
		}
	}
}

// main function of the bulk directory management program.
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

		batchProcess(mode, batchSize, dirs...)
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

		batchProcess("delete", batchSize, dirs...)

	default:
		log.Fatalf("Invalid mode '%s'. Mode must be 'create' or 'delete'.", mode)
	}

}
