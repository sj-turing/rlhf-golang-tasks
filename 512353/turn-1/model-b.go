package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/mail"
	"os"
	"strings"
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
	data, err := ioutil.ReadFile(filePath)
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

func main() {
	// Create a context with a timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Read users from a JSON file
	filePath := "users.json"
	users, err := ReadUsersFromJSON(ctx, filePath)
	if err != nil {
		log.Fatalf("Error reading users from JSON file: %v", err)
	}

	// Filter valid users
	validUsers, err := FilterValidUsers(ctx, users)
	if err != nil {
		log.Fatalf("Error filtering valid users: %v", err)
	}

	// Print the valid users
	for _, user := range validUsers {
		fmt.Printf("Username: %s, Spend Amount: %.2f, Email: %s\n", user.Username, user.SpendAmount, user.Email)
	}
}
