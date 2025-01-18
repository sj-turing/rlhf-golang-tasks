// filename: errors.go
package errors

import (
	"encoding/json"
	"net/http"
)

// bankError custom type helps to write possible errors for banking applications
type BankError uint16

const (
	ErrNegativeAmountNotAllowed BankError = iota + 1
	ErrInsufficientBalance
	ErrInvalidAccountID
	ErrInternalServerError // Generic error for unexpected conditions
	ErrMethodNotAllow
	ErrInvalidRequestPayload
)

// ErrMessageByBankError maps meaningful message with bank error
var errMessageByBankError = map[BankError]string{
	ErrNegativeAmountNotAllowed: "negative amount not allowed",
	ErrInsufficientBalance:      "insufficient balance",
	ErrInvalidAccountID:         "invalid accountID",
	ErrInternalServerError:      "internal server error",
	ErrMethodNotAllow:           "method not implemented",
	ErrInvalidRequestPayload:    "invalid request payload",
}

// ErrCodeByBankError maps HTTP status codes with bank errors
var errCodeByBankError = map[BankError]int{
	ErrNegativeAmountNotAllowed: http.StatusPreconditionFailed,  // 412
	ErrInsufficientBalance:      http.StatusBadRequest,          // 400
	ErrInvalidAccountID:         http.StatusBadRequest,          // 400
	ErrInternalServerError:      http.StatusInternalServerError, // 500
	ErrMethodNotAllow:           http.StatusMethodNotAllowed,
	ErrInvalidRequestPayload:    http.StatusBadRequest,
}

// errorResponse helps to format error message for readability
type errorResponse struct {
	Message string `json:"error"`
	Code    int    `json:"code"`
}

// Error, returns custom format for the error
func (be BankError) Error() string {
	msg, ok := errMessageByBankError[be]
	errResp := errorResponse{}
	if !ok {
		errResp.Message = "error not implemented"
		errResp.Code = http.StatusInternalServerError
	} else {
		errResp.Message = msg
		errResp.Code = be.GetStatusCode()
	}
	bb, _ := json.Marshal(errResp)
	return string(bb)
}

// GetStatusCode returns the HTTP status code associated with the error
func (be BankError) GetStatusCode() int {
	code, ok := errCodeByBankError[be]
	if !ok {
		return http.StatusInternalServerError // Default to 500 if not mapped
	}
	return code
}

// GetMessage returns the meaningful message associated with the error
func (be BankError) GetMessage() string {
	message, ok := errMessageByBankError[be]
	if !ok {
		return "Internal Server Error" // Default to 500 if not mapped
	}
	return message
}
