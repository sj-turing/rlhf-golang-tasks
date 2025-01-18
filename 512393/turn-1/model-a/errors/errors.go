package errors

import (
	"fmt"
)

// BaseError serves as a base type for all custom errors, embedding the Go error interface.
type BaseError struct {
	Operation string
	Err       error
}

func (e *BaseError) Error() string {
	return fmt.Sprintf("error during %s: %v", e.Operation, e.Err)
}

// DebitError represents errors specific to the Debit API.
type DebitError struct {
	BaseError
	AccountID string
}

// CreditError represents errors specific to the Credit API.
type CreditError struct {
	BaseError
	AccountID string
}

// BalanceCheckError represents errors specific to the Balance Check API.
type BalanceCheckError struct {
	BaseError
	AccountID string
}

func WrapDebitError(accountID string, err error) error {
	return &DebitError{
		BaseError: BaseError{
			Operation: "debit",
			Err:       err,
		},
		AccountID: accountID,
	}
}

func WrapCreditError(accountID string, err error) error {
	return &CreditError{
		BaseError: BaseError{
			Operation: "credit",
			Err:       err,
		},
		AccountID: accountID,
	}
}

func WrapBalanceCheckError(accountID string, err error) error {
	return &BalanceCheckError{
		BaseError: BaseError{
			Operation: "balance check",
			Err:       err,
		},
		AccountID: accountID,
	}
}
