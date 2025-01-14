package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var batchSize = 10
var ctx = context.Background()
var cancel context.CancelFunc

func init() {
	ctx, cancel = context.WithCancel(ctx)
	// Set up signal handler to cancel operations on interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for {
			<-c
			cancel()
			break
		}
	}()
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter directories to create (empty line to finish):")
	var dirsToCreate []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		line = strings.TrimSpace(line)
		if line != "" {
			dirsToCreate = append(dirsToCreate, line)
		}
	}

	createDirectories(dirsToCreate)

	fmt.Println("Enter directories to delete (empty line to finish):")
	var dirsToDelete []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		line = strings.TrimSpace(line)
		if line != "" {
			dirsToDelete = append(dirsToDelete, line)
		}
	}

	deleteDirectories(dirsToDelete)
}

func createDirectories(dirs []string) {
	for i := 0; i < len(dirs); i += batchSize {
		batch := dirs[i : i+batchSize]
		createBatch(ctx, batch)
	}
}

func createBatch(ctx context.Context, dirs []string) {
	wg := &sync.WaitGroup{}

	for _, dir := range dirs {
		wg.Add(1)
		go func(dir string) {
			defer wg.Done()

			if err := os.MkdirAll(dir, 0755); err != nil {
				log.Printf("Failed to create %s: %v", dir, err)
			}
		}(dir)
	}
	wg.Wait()
}

func deleteDirectories(dirs []string) {
	for i := 0; i < len(dirs); i += batchSize {
		batch := dirs[i : i+batchSize]
		deleteBatch(ctx, batch)
	}
}

func deleteBatch(ctx context.Context, dirs []string) {
	wg := &sync.WaitGroup{}

	for _, dir := range dirs {
		wg.Add(1)
		go func(dir string) {
			defer wg.Done()

			err := os.RemoveAll(dir)
			if err != nil {
				// Check for specific errors and handle accordingly
				if os.IsNotExist(err) {
					log.Printf("Directory %s does not exist.", dir)
				} else {
					log.Printf("Failed to delete %s: %v", dir, err)
				}
			}
		}(dir)
	}
	wg.Wait()
}
