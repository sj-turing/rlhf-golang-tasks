package main

import (
	customErrors "512393/turn-1/model-a/errors"
	"errors"
	"fmt"
)

func debit(accountID string, amount float64) error {
	// Simulate an error
	if amount < 0 {
		return customErrors.WrapDebitError(accountID, errors.New("negative amount not allowed"))
	}
	// Logic for debiting the account...
	return nil
}

func main() {
	err := debit("12345", -100.0)
	if err != nil {
		// Custom handling for specific error types
		switch e := err.(type) {
		case *customErrors.DebitError:
			fmt.Printf("Failed to debit account %s: %v\n", e.AccountID, e)
		default:
			fmt.Println("An unexpected error occurred:", err)
		}
	}
}
