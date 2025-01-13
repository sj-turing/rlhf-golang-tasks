package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// Define a struct to hold the query parameters
type AccountDetailsRequest struct {
	UserID     string `validate:"required,min=6"`
	DeviceID   string `validate:"required,min=6"`
	BankCode   string `validate:"required,min=2,max=4"`
	MerchantID string `validate:"required,min=6"`
}

// Validate the request parameters
func validateAccountDetailsRequest(r *http.Request) (*AccountDetailsRequest, error) {
	v := validator.New()
	var req AccountDetailsRequest

	// Parse query parameters
	err := r.ParseForm()
	if err != nil {
		return nil, err
	}

	// Bind the query parameters to the struct
	err = v.Struct(&req)
	if err != nil {
		return nil, err
	}

	// Output validated fields
	fmt.Println("UserID:", req.UserID)
	fmt.Println("DeviceID:", req.DeviceID)
	fmt.Println("BankCode:", req.BankCode)
	fmt.Println("MerchantID:", req.MerchantID)

	return &req, nil
}

// Example HTTP handler for the GetAccountDetails API
func getAccountDetailsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	req, err := validateAccountDetailsRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Success response
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Query parameters are valid. Proceeding with request.")
}

// Start the server
func main() {
	http.HandleFunc("/get-account-details", getAccountDetailsHandler)
	fmt.Println("Server running on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
