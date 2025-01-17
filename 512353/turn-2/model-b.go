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

func ValidateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func processFile(ctx context.Context, filePath string, spendAmountChan chan<- float64, wg *sync.WaitGroup) {
	defer wg.Done()
	// Check if the context has been cancelled or timed out
	if ctx.Err() != nil {
		return
	}
	// Read the JSON file
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("Error reading file %s: %v\n", filePath, err)
		return
	}
	// Decode the JSON data into a slice of User structs
	var users []User
	if err := json.Unmarshal(data, &users); err != nil {
		log.Printf("Error decoding JSON from file %s: %v\n", filePath, err)
		return
	}
	// Iterate through the slice of users and validate each user's email address
	for _, user := range users {
		// Check if the context has been cancelled or timed out
		if ctx.Err() != nil {
			return
		}
		if ValidateEmail(user.Email) {
			spendAmountChan <- user.SpendAmount
		}
	}
}

func calculateTotalSpend(ctx context.Context, spendAmountChan <-chan float64, wg *sync.WaitGroup) {
	defer wg.Done()
	totalSpend := 0.0
	for {
		select {
		case amount := <-spendAmountChan:
			totalSpend += amount
		case <-ctx.Done():
			log.Println("Calculation goroutine canceled.")
			return
		}
	}
	fmt.Println("Total Spend:", totalSpend)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // Increase timeout as needed
	defer cancel()

	numFiles := 100
	spendAmountChan := make(chan float64, numFiles) // Buffered channel to handle concurrent sends
	var wg sync.WaitGroup

	// Start the calculation goroutine
	wg.Add(1)
	go calculateTotalSpend(ctx, spendAmountChan, &wg)

	// Start goroutines to process each JSON file concurrently
	for i := 1; i <= numFiles; i++ {
		filePath := fmt.Sprintf("./json_files/users_%d.json", i)
		wg.Add(1)
		go processFile(ctx, filePath, spendAmountChan, &wg)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	close(spendAmountChan)
}
