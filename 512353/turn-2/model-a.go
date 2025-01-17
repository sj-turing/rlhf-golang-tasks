package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/mail"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// User represents a user's data
type User struct {
	Username    string  `json:"username"`
	SpendAmount float64 `json:"spendAmount"`
	Email       string  `json:"email"`
}

// ValidateEmail validates an email address
func ValidateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// ReadUsersFromJSON reads users from a JSON file
func ReadUsersFromJSON(ctx context.Context, filePath string) ([]User, error) {
	// Check if the context has been cancelled or timed out
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	// Read the JSON file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Decode the JSON data into a slice of User structs
	var users []User
	err = json.Unmarshal(data, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// FilterValidUsers filters valid users from a slice of User structs
func FilterValidUsers(ctx context.Context, users []User) ([]User, error) {
	// Check if the context has been cancelled or timed out
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	var validUsers []User

	// Iterate through the slice of users and validate each user's email address
	for _, user := range users {
		// Check if the context has been cancelled or timed out
		if ctx.Err() != nil {
			return nil, ctx.Err()
		}

		if ValidateEmail(user.Email) {
			validUsers = append(validUsers, user)
		}
	}

	return validUsers, nil
}

// ProcessFile reads and processes a single JSON file
func ProcessFile(ctx context.Context, filePath string, results chan<- float64, wg *sync.WaitGroup) {
	defer wg.Done()

	users, err := ReadUsersFromJSON(ctx, filePath)
	if err != nil {
		log.Printf("Error reading users from file %s: %v", filePath, err)
		return
	}

	validUsers, err := FilterValidUsers(ctx, users)
	if err != nil {
		log.Printf("Error filtering valid users from file %s: %v", filePath, err)
		return
	}

	var totalSpend float64
	for _, user := range validUsers {
		totalSpend += user.SpendAmount
	}

	select {
	case <-ctx.Done():
		log.Printf("Context done for file %s: %v", filePath, ctx.Err())
	case results <- totalSpend:
	}
}

func main() {
	// Create a context with a timeout of 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Directory containing the JSON files
	dir := "./json_files"

	// Channel for receiving spend amounts
	results := make(chan float64, 100)

	// WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Walk through the directory and process each file
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".json" {
			wg.Add(1)
			go ProcessFile(ctx, path, results, &wg)
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Error walking the path %s: %v", dir, err)
	}

	// Close the results channel when all goroutines have finished
	go func() {
		wg.Wait()
		close(results)
	}()

	// Aggregate the total spendAmount
	var totalSpendAmount float64
	for spend := range results {
		totalSpendAmount += spend
	}

	fmt.Printf("Total Spend Amount: %.2f\n", totalSpendAmount)
}
