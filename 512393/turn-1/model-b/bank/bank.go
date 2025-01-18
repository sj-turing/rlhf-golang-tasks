package bank

import "fmt"

// Custom error types for banking operations

type InvalidAmountError struct {
	Amount float64
	Reason string
}

func (e *InvalidAmountError) Error() string {
	return fmt.Sprintf("Invalid amount: %.2f - %s", e.Amount, e.Reason)
}

type InsufficientBalanceError struct {
	Balance float64
	Min     float64
}

func (e *InsufficientBalanceError) Error() string {
	return fmt.Sprintf("Insufficient balance: %.2f (min: %.2f)", e.Balance, e.Min)
}

type OverdraftLimitError struct {
	Amount float64
	Limit  float64
}

func (e *OverdraftLimitError) Error() string {
	return fmt.Sprintf("Overdraft limit exceeded: %.2f (limit: %.2f)", e.Amount, e.Limit)
}

func Debit(amount float64) error {
	if amount <= 0 {
		return &InvalidAmountError{Amount: amount, Reason: "Must be positive"}
	}
	// Debit transaction logic here
	return nil
}

func Credit(amount float64) error {
	if amount <= 0 {
		return &InvalidAmountError{Amount: amount, Reason: "Must be positive"}
	}
	// Credit transaction logic here
	return nil
}

func CheckBalance() (float64, error) {
	// Balance check logic here
	return 0, nil
}

func ProcessTransaction(debitAmount float64, creditAmount float64) (float64, error) {
	err := Debit(debitAmount)
	if err != nil {
		return 0, err
	}
	err = Credit(creditAmount)
	if err != nil {
		return 0, err
	}
	balance, err := CheckBalance()
	if err != nil {
		return 0, err
	}
	return balance, nil
}

func ProcessTransaction(debitAmount float64, creditAmount float64) (float64, error) {
	if err := Debit(debitAmount); err != nil {
		return 0, err
	}
	if err := Credit(creditAmount); err != nil {
		return 0, err
	}
	balance, err := CheckBalance()
	if err != nil {
		return 0, err
	}
	return balance, nil
}
