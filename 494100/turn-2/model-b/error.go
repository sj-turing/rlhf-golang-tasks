package bankerrors

import "fmt"

type BankError struct {
	Code    int
	Message string
}

func (be *BankError) Error() string {
	return fmt.Sprintf("Error code: %d, Message: %s", be.Code, be.Message)
}

const (
	ErrInsufficientBalanceCode  = 2001
	ErrInvalidAccountNumberCode = 2002
	ErrAccountNotFoundCode      = 2003
	ErrInternalServerErrorCode  = 2004
)

var errorMap = map[int]string{
	ErrInsufficientBalanceCode:  "Insufficient balance in account",
	ErrInvalidAccountNumberCode: "Invalid account number format",
	ErrAccountNotFoundCode:      "Account not found",
	ErrInternalServerErrorCode:  "Internal server error occurred",
}

func NewInsufficientBalanceError() error {
	return &BankError{Code: ErrInsufficientBalanceCode, Message: errorMap[ErrInsufficientBalanceCode]}
}
func NewInvalidAccountNumberError() error {
	return &BankError{Code: ErrInvalidAccountNumberCode, Message: errorMap[ErrInvalidAccountNumberCode]}
}
func NewAccountNotFoundError() error {
	return &BankError{Code: ErrAccountNotFoundCode, Message: errorMap[ErrAccountNotFoundCode]}
}
func NewInternalServerError() error {
	return &BankError{Code: ErrInternalServerErrorCode, Message: errorMap[ErrInternalServerErrorCode]}
}
