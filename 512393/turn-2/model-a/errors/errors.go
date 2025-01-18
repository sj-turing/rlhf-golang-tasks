// filename: errors.go
package errors

import (
	"encoding/json"
	"net/http"
)

// bankError custom type helps to write possible errors for banking applications
type bankError uint16

const (
	ErrNegativeAmountNotAllowed bankError = iota + 1
	ErrInsufficientBalance
	ErrInvalidAccountID
	ErrInternalServerError // Generic error for unexpected conditions
)

// ErrMessageByBankError maps meaningful message with bank error
var ErrMessageByBankError = map[bankError]string{
	ErrNegativeAmountNotAllowed: "negative amount not allowed",
	ErrInsufficientBalance:      "insufficient balance",
	ErrInvalidAccountID:         "invalid accountID",
	ErrInternalServerError:      "internal server error",
}

// ErrCodeByBankError maps HTTP status codes with bank errors
var ErrCodeByBankError = map[bankError]int{
	ErrNegativeAmountNotAllowed: http.StatusPreconditionFailed,  // 412
	ErrInsufficientBalance:      http.StatusBadRequest,          // 400
	ErrInvalidAccountID:         http.StatusBadRequest,          // 400
	ErrInternalServerError:      http.StatusInternalServerError, // 500
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
		errResp.Message = "error not implemented"
	} else {
		errResp.Message = msg
	}
	bb, _ := json.Marshal(errResp)
	return string(bb)
}

// StatusCode returns the HTTP status code associated with the error
func (be bankError) StatusCode() int {
	code, ok := ErrCodeByBankError[be]
	if !ok {
		return http.StatusInternalServerError // Default to 500 if not mapped
	}
	return code
}
