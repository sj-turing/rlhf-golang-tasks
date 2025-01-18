// filename: main.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"512393/turn-2/model-a/errors" // path to import above mentioned error package
)

// AccountDetails holds account information
type AccountDetails struct {
	AccountID string  `json:"account_id"`
	Balance   float32 `json:"balance"`
}

// In-memory account data
var accounts = map[string]float32{
	"acc123": 1000.0,
	"acc456": 500.0,
}

func main() {
	http.HandleFunc("/debit", debitHandler)
	http.HandleFunc("/credit", creditHandler)
	http.HandleFunc("/balance", balanceHandler)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}

func debitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var account AccountDetails
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := validateAccountID(account.AccountID); err != nil {
		respondWithError(w, err)
		return
	}

	if err := validateBalance(account.Balance); err != nil {
		respondWithError(w, err)
		return
	}

	// Perform debit operation
	if balance, exists := accounts[account.AccountID]; exists {
		if balance < account.Balance {
			respondWithError(w, errors.ErrInsufficientBalance)
			return
		}
		accounts[account.AccountID] -= account.Balance
		respondWithJSON(w, http.StatusOK, map[string]float32{"new_balance": accounts[account.AccountID]})
	} else {
		respondWithError(w, errors.ErrInvalidAccountID)
	}
}

func creditHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var account AccountDetails
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := validateAccountID(account.AccountID); err != nil {
		respondWithError(w, err)
		return
	}

	if err := validateBalance(account.Balance); err != nil {
		respondWithError(w, err)
		return
	}

	// Perform credit operation
	if _, exists := accounts[account.AccountID]; exists {
		accounts[account.AccountID] += account.Balance
		respondWithJSON(w, http.StatusOK, map[string]float32{"new_balance": accounts[account.AccountID]})
	} else {
		respondWithError(w, errors.ErrInvalidAccountID)
	}
}

func balanceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	accountID := r.URL.Query().Get("account_id")
	if err := validateAccountID(accountID); err != nil {
		respondWithError(w, err)
		return
	}

	// Perform balance check operation
	if balance, exists := accounts[accountID]; exists {
		respondWithJSON(w, http.StatusOK, map[string]float32{"balance": balance})
	} else {
		respondWithError(w, errors.ErrInvalidAccountID)
	}
}

func respondWithError(w http.ResponseWriter, err error) {
	var be errors.bankError
	if ok := interface{}(err).(errors.bankError); ok {
		http.Error(w, be.Error(), be.StatusCode())
	} else {
		http.Error(w, errors.ErrInternalServerError.Error(), http.StatusInternalServerError)
	}
}

func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func validateBalance(balance float32) error {
	if balance < 0 {
		return errors.ErrNegativeAmountNotAllowed
	}

	return nil
}

func validateAccountID(accountID string) error {
	if accountID == "" {
		return errors.ErrInvalidAccountID
	}

	return nil
}
