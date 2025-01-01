package main

import (
	"fmt"
	"net/http"

	"494100/turn-2/model-a/errors/banking"
	"494100/turn-2/model-a/errors/services/banking"
)

func main() {
	http.HandleFunc("/api/account/create", createAccountHandler)
	http.HandleFunc("/api/account/deposit", depositHandler)
	http.HandleFunc("/api/account/withdraw", withdrawHandler)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

func createAccountHandler(w http.ResponseWriter, r *http.Request) {
	owner := r.URL.Query().Get("owner")
	account, err := banking.CreateAccount(owner)
	if err != nil {
		handleError(w, err)
		return
	}

	respondJSON(w, http.StatusCreated, account)
}

func depositHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	amountStr := r.URL.Query().Get("amount")
	amount, err := fmt.ParseFloat(amountStr, 64)
	if err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	amountDeposited, err := banking.Deposit(id, amount)
	if err != nil {
		handleError(w, err)
		return
	}

	respondJSON(w, http.StatusOK, map[string]interface{}{"deposited": amountDeposited})
}

func withdrawHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	amountStr := r.URL.Query().Get("amount")
	amount, err := fmt.ParseFloat(amountStr, 64)
	if err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	amountWithdrawn, err := banking.Withdraw(id, amount)
	if err != nil {
		handleError(w, err)
		return
	}

	respondJSON(w, http.StatusOK, map[string]interface{}{"withdrawn": amountWithdrawn})
}

func handleError(w http.ResponseWriter, err error) {
	switch {
	case banking.ErrInvalidOwner == err:
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
	case banking.ErrDepositAmountInvalid == err || banking.ErrWithdrawalAmountInvalid == err:
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
	case banking.ErrAccountNotFound == err:
		respondJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
	case banking.ErrInsufficientBalance == err:
		respondJSON(w, http.StatusForbidden, map[string]string{"error": err.Error()})
	default:
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
	}
}

func respondJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(data)
}
