package main

/*

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Cancel the context at the end of the main function

	go work(ctx)

	select {
	case <-time.After(2 * time.Second): // Simulate waiting for the goroutine
		log.Println("Canceling work due to timeout")
		cancel() // Cancel the work if it's taking too long
	}
}

func work(ctx context.Context) {
	file, err := os.Create("example.txt")
	if err != nil {
		handleError(err)
		return
	}
	defer file.Close() // Ensure the file is closed regardless of the exit path

	for {
		select {
		case <-ctx.Done(): // Check for context cancellation
			log.Println("Work canceled")
			return
		default:
			_, err := file.WriteString("Hello, world!\n")
			if err != nil {
				handleError(err)
				return
			}
			fmt.Println("Wrote to file")
			time.Sleep(1 * time.Second)
		}
	}
}

func handleError(err error) {
	if err != nil {
		log.Printf("Error: %v\n", err)
		panic(err) // Typically, you wouldn't panic in production code, but using it here for illustration
	}
}

*/
