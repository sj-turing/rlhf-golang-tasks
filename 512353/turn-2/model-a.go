package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/mail"
	"os"
	"sync"
	"time"
)

type User struct {
	Username    string  `json:"username"`
	SpendAmount float64 `json:"spendAmount"`
	Email       string  `json:"email"`
}

func validateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func readUsersFromJSON(ctx context.Context, filePath string) ([]User, error) {
	// Do not proceed if the context is canceled
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var users []User
	err = json.Unmarshal(data, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func processFile(ctx context.Context, filePath string, resultChan chan<- float64, wg *sync.WaitGroup) {
	defer wg.Done() // Signal that this goroutine is done

	users, err := readUsersFromJSON(ctx, filePath)
	if err != nil {
		log.Printf("Failed to read file %s: %v", filePath, err)
		return
	}

	for _, user := range users {
		if ctx.Err() != nil {
			return // Stop processing if context is canceled
		}
		if validateEmail(user.Email) {
			select {
			case resultChan <- user.SpendAmount:
			case <-ctx.Done():
				return // Stop if context is done
			}
		}
	}
}

func main() {
	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filePaths := make([]string, 100) // Assume these are the paths to 100 JSON files
	for i := 0; i < 100; i++ {
		filePaths[i] = fmt.Sprintf("./json_files/user%d.json", i+1) // Example filenames
	}

	resultChan := make(chan float64)
	var wg sync.WaitGroup

	// Summation goroutine
	go func() {
		totalSpend := 0.0
		for spendAmount := range resultChan {
			totalSpend += spendAmount
		}
		fmt.Printf("Total spend amount: %.2f\n", totalSpend)
	}()

	// Launch a goroutine for each file
	for _, filePath := range filePaths {
		wg.Add(1)
		go processFile(ctx, filePath, resultChan, &wg)
	}

	// Wait for all file processing goroutines to complete, then close the channel
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Wait for all processing to complete
	select {
	case <-ctx.Done():
		log.Println("Operation timed out or was canceled.")
	}
}
