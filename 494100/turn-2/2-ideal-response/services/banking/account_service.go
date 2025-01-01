package banking_service

import (
	"494100/turn-2/2-ideal-response/errors/banking"
	"494100/turn-2/2-ideal-response/models"
)

// CreateAccount creates a new account
func CreateAccount(owner string) (models.Account, error) {
	account := models.Account{
		Owner: owner,
	}

	// Simulate an error when the owner is invalid
	if owner == "" {
		return account, banking.ErrInvalidOwner
	}

	return account, nil
}

// Deposits funds into an account
func Deposit(id string, amount float64) (float64, error) {
	if amount <= 0 {
		return 0, banking.ErrDepositAmountInvalid
	}

	// Simulate an error when account doesn't exist
	if id == "" {
		return 0, banking.ErrAccountNotFound
	}

	return amount, nil
}

// Withdraws funds from an account
func Withdraw(id string, amount float64) (float64, error) {
	if amount <= 0 {
		return 0, banking.ErrWithdrawalAmountInvalid
	}

	// Simulate an error when account doesn't exist
	if id == "" {
		return 0, banking.ErrAccountNotFound
	}

	// Simulate an insufficient balance error
	if amount > 100 { // Hardcoded limit for simplicity
		return 0, banking.ErrInsufficientBalance
	}

	return amount, nil
}
