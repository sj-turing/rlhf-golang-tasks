package bankapi

import (
	"bankerrors"
	"fmt"
	"net/http"
)

type Account struct {
	Number  string
	Balance float64
}

func Deposit(accountNumber string, amount float64) error {
	// Simulate business logic to check account balance
	if amount <= 0 {
		return bankerrors.NewInvalidAccountNumberError()
	}
	return nil
}
func HandleDepositRequest(w http.ResponseWriter, r *http.Request) {
	// Simulate request processing
	// ...
	// ...

	// Depositing the amount
	err := Deposit(accountNumber, amount)

	// Handle errors using the tailored error package
	if err != nil {
		handleBankError(w, err)
		return
	}

	fmt.Fprintf(w, "Deposit successful! New balance: %f", updatedBalance)
}
func handleBankError(w http.ResponseWriter, err error) {
	be, ok := err.(*bankerrors.BankError)
	if !ok {
		be = bankerrors.NewInternalServerError()
	}

	// Set the appropriate HTTP status code based on the error code
	statusCode := http.StatusInternalServerError
	switch be.Code {
	case bankerrors.ErrInsufficientBalanceCode:
		statusCode = http.StatusUnprocessableEntity
	case bankerrors.ErrInvalidAccountNumberCode:
		statusCode = http.StatusBadRequest
	case bankerrors.ErrAccountNotFoundCode:
		statusCode = http.StatusNotFound
	}

	w.WriteHeader(statusCode)
	fmt.Fprintf(w, "Error: %s\n", be.Message)
}
