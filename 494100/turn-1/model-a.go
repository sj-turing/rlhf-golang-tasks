package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Error codes
const (
	ErrInvalidInput        = 100
	ErrResourceNotFound    = 200
	ErrInternalServerError = 500
)

// Custom error types
type InputError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (e InputError) Error() string {
	return e.Message
}

type ResourceError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (e ResourceError) Error() string {
	return e.Message
}

type ServerError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (e ServerError) Error() string {
	return e.Message
}

func processRequest(input string) error {
	// Simulate business logic validation
	if input == "" {
		return InputError{Message: "Input is required", Code: ErrInvalidInput}
	}

	// Simulate resource lookup
	if input == "missing" {
		return ResourceError{Message: "Resource not found", Code: ErrResourceNotFound}
	}

	return nil
}

func handleError(err error, w http.ResponseWriter, r *http.Request) {
	switch err {
	case InputError{}:
		response := InputError{
			Message: err.Error(),
			Code:    ErrInvalidInput,
		}
		http.Error(w, json.Marshal(response), http.StatusBadRequest)
	case ResourceError{}:
		response := ResourceError{
			Message: err.Error(),
			Code:    ErrResourceNotFound,
		}
		http.Error(w, json.Marshal(response), http.StatusNotFound)
	case ServerError{}:
		response := ServerError{
			Message: err.Error(),
			Code:    ErrInternalServerError,
		}
		http.Error(w, json.Marshal(response), http.StatusInternalServerError)
	default:
		response := ServerError{
			Message: "An unknown error occurred",
			Code:    ErrInternalServerError,
		}
		http.Error(w, json.Marshal(response), http.StatusInternalServerError)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	input := r.URL.Query().Get("input")

	err := processRequest(input)
	if err != nil {
		handleError(err, w, r)
		return
	}

	// Simulate successful business logic
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Processed successfully")
}

func main() {
	http.HandleFunc("/process", handler)

	log.Println("Service starting on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
