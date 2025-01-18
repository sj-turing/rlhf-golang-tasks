package errors

import (
	"encoding/json"
)

// bankError custom type helps to write possible errors for banking applications
type bankError uint16

const (
	ErrNegativeAmountNotAllowed bankError = iota + 1
	ErrInsufficientBalance
	ErrInvalidAccountID
	// all other errors can go here
)

// ErrMessageByBankError maps meaningful message with bank error
var ErrMessageByBankError = map[bankError]string{
	ErrNegativeAmountNotAllowed: "negative amount not allowed",
	ErrInsufficientBalance:      "insufficient balance",
	ErrInvalidAccountID:         "invalid accountID",
	// add meaningful message to the custom error
}

// errorResponse helps to format error message for readability
type errorResponse struct {
	Message string `json:"error"`
}

// Error, returns custom format for the error
func (be bankError) Error() string {

	msg, ok := ErrMessageByBankError[be]
	errResp := errorResponse{}
	if !ok {
		// if by mistakely forgot to map error with meaningful message
		// adding 'error not implemented' message
		errResp.Message = "error not implemented"
	}
	errResp.Message = msg
	bb, _ := json.Marshal(errResp)
	return string(bb)
}
