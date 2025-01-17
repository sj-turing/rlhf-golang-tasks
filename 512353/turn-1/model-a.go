package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/mail"
	"os"
	"time"
)

type User struct {
	Username    string  `json:"username"`
	SpendAmount float64 `json:"spendAmount"`
	Email       string  `json:"email"`
}

func main() {
	// Set a timeout for the entire operation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	users, err := readAndFilterUsers(ctx, "users.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, user := range users {
		fmt.Printf("Valid user: %v\n", user)
	}
}

func readAndFilterUsers(ctx context.Context, filename string) ([]User, error) {
	data, err := readFileWithContext(ctx, filename)
	if err != nil {
		return nil, err
	}

	var users []User
	if err := json.Unmarshal(data, &users); err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	var validUsers []User
	for _, user := range users {
		if isValidEmail(user.Email) {
			validUsers = append(validUsers, user)
		}
	}

	return validUsers, nil
}

func readFileWithContext(ctx context.Context, filename string) ([]byte, error) {
	done := make(chan error, 1)
	var data []byte

	go func() {
		var err error
		data, err = ioutil.ReadFile(filename)
		done <- err
	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err() // Context canceled or timed out
	case err := <-done:
		if err != nil {
			return nil, fmt.Errorf("error reading file: %w", err)
		}
	}

	return data, nil
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
