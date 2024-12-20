package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

// SuccessResponse is a custom successful response
type SuccessResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ErrorResponse is a custom error response
type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// DownstreamServiceError is a custom error type with annotations
type DownstreamServiceError struct {
	Err         error  // The underlying error
	Service     string // The name of the downstream service
	Method      string // The method called on the service
	Description string // Additional description of the failure
}

// Error implements the error interface
func (e DownstreamServiceError) Error() string {
	return fmt.Sprintf("downstream service error: service=%s, method=%s, description=%s, error=%v", e.Service, e.Method, e.Description, e.Err)
}

// CallDownstreamService is a function that simulates calling a downstream service
func CallDownstreamService(service, method string) error {
	// Simulate a network error
	if service == "PaymentService" {
		return DownstreamServiceError{Err: errors.New("network error"), Service: service, Method: method, Description: "Failed to connect to PaymentService"}
	}
	// Simulate a service-specific error
	if service == "AuthenticationService" && method == "Authenticate" {
		return DownstreamServiceError{Err: errors.New("invalid credentials"), Service: service, Method: method, Description: "Authentication failed"}
	}
	// Simulate a business logic error
	if service == "InventoryService" && method == "CheckStock" {
		return DownstreamServiceError{Err: errors.New("item out of stock"), Service: service, Method: method, Description: "Item is out of stock"}
	}
	return nil
}

// handleSomeOperation is the HTTP handler for performing some operation
func handleSomeOperation(w http.ResponseWriter, r *http.Request) {
	// Extract necessary data from the request
	service := r.URL.Query().Get("service")
	method := r.URL.Query().Get("method")

	if service == "" || method == "" {
		http.Error(w, "service and method are required", http.StatusBadRequest)
		return
	}

	err := CallDownstreamService(service, method)
	if err != nil {
		// Handle downstream service error
		switch err.(type) {
		case DownstreamServiceError:
			dse := err.(DownstreamServiceError)
			byteError, err := json.Marshal(ErrorResponse{Status: "error", Message: dse.Description, Code: http.StatusInternalServerError})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			http.Error(w, string(byteError), http.StatusInternalServerError)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Handle successful operation
	successResponse := SuccessResponse{Status: "success", Message: "Operation successful", Data: map[string]string{"service": service, "method": method}}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(successResponse); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/someOperation", handleSomeOperation)

	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
