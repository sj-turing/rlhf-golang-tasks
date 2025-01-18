package main

import (
	"fmt"

	"512393/turn-1/1-ideal-response/errors"
)

// AccountDetails structs to hold account detail of the user
type AccountDetails struct {
	AccountID string
	Balance   float32
}

func main() {
	debit1 := AccountDetails{AccountID: "abc-1", Balance: 0}
	if err := validateBalance(debit1.Balance); err != nil {
		fmt.Println(err)
	}

	debit2 := AccountDetails{AccountID: "pqr-1", Balance: -19}
	if err := validateBalance(debit2.Balance); err != nil {
		fmt.Println(err)
	}

	credit := AccountDetails{AccountID: "", Balance: 123}
	if err := validateAccountID(credit.AccountID); err != nil {
		fmt.Println(err)
	}
}

// validateBalance returns an error if balance is invalid
func validateBalance(balance float32) error {
	if balance == 0 {
		return errors.ErrInsufficientBalance
	}

	if balance < 0 {
		return errors.ErrNegativeAmountNotAllowed
	}

	return nil
}

// validateAccountID returns an error if accountID is invalid
func validateAccountID(accountID string) error {
	if accountID == "" {
		return errors.ErrInvalidAccountID
	}

	return nil
}
