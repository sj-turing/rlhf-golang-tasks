// filename: main.go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strconv"

    "errors" // path to import above mentioned error package
)

// AccountDetails structs to hold account detail of the user
type AccountDetails struct {
    AccountID string
    Balance   float32
}

func main() {
    http.HandleFunc("/debit", debitHandler)
    http.HandleFunc("/credit", creditHandler)
    http.HandleFunc("/balance", balanceHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func debitHandler(w http.ResponseWriter, r *http.Request) {
    // Implement debit logic here
    // For demonstration purposes, let's simulate a debit operation
    accountID := r.FormValue("accountID")
    amountStr := r.FormValue("amount")
    amount, err := strconv.ParseFloat(amountStr, 32)
    if err != nil {
        writeErrorResponse(w, http.StatusBadRequest, errors.ErrInvalidAmountFormat)
        return
    }

    // Perform the debit operation and handle errors
    err = performDebit(accountID, amount)
    if err != nil {
        writeErrorResponse(w, http.StatusInternalServerError, err)
        return
    }

    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Debit successful")
}

func creditHandler(w http.ResponseWriter, r *http.Request) {
    // Implement credit logic here
    // For demonstration purposes, let's simulate a credit operation
    accountID := r.FormValue("accountID")
    amountStr := r.FormValue("amount")
    amount, err := strconv.ParseFloat(amountStr, 32)
    if err != nil {
        writeErrorResponse(w, http.StatusBadRequest, errors.ErrInvalidAmountFormat)
        return
    }

    // Perform the credit operation and handle errors
    err = performCredit(accountID, amount)
    if err != nil {
        writeErrorResponse(w, http.StatusInternalServerError, err)
        return
    }

    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Credit successful")
}

func balanceHandler(w http.ResponseWriter, r *http.Request) {
    // Implement balance check logic here
    // For demonstration purposes, let's simulate a balance check operation
    accountID := r.FormValue("accountID")

    balance, err := checkBalance(accountID)
    if err != nil {
        writeErrorResponse(w, http.StatusInternalServerError, err)
        return
    }

    resp := struct {
        Balance float32 `json:"balance"`
    }{
        Balance: balance,
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(resp)
}

func performDebit(accountID string, amount float32) error {
    if amount < 0 {
        return errors.ErrNegativeAmountNotAllowed
    }
    // Add more debit validation logic here
    return nil
}

func performCredit(accountID string, amount float32) error {
    if amount < 0 {
        return errors.ErrNegativeAmountNotAllowed
    }
    // Add more credit validation logic here
    return nil
}

func checkBalance(accountID string) (float32, error) {
    // Add balance check logic here
    // For demonstration, let's return a dummy balance
    return 100.0, nil
}

func writeErrorResponse(w http.ResponseWriter, statusCode int, err error) {
    var httpStatusCode int
    var errMessage string

    // Use type assertion to extract the underlying bankError if available
    if be, ok := err.(errors.bankError); ok {
        httpStatusCode = int(be)
        errMessage = be.Error()
    } else {
