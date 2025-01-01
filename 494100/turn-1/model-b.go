package main

import (
	"fmt"
	"net/http"
)

type CustomError struct {
	Code    int
	Message string
}

func (ce *CustomError) Error() string {
	return fmt.Sprintf("Error code: %d, Message: %s", ce.Code, ce.Message)
}

const (
	ErrInvalidRequestCode      = 1001
	ErrResourceNotFoundCode    = 1002
	ErrInternalServerErrorCode = 1003
)

var errorMap = map[int]string{
	ErrInvalidRequestCode:      "Invalid request parameters",
	ErrResourceNotFoundCode:    "Resource not found",
	ErrInternalServerErrorCode: "Internal server error",
}

func main() {
	http.HandleFunc("/", homeHandler)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Sample business logic that can trigger errors
	// ...
	// ...

	// Simulate an error condition
	err := &CustomError{Code: ErrInvalidRequestCode, Message: "Missing required parameter 'name'"}

	// Handle error using the error map
	handleError(w, err)
}
func handleError(w http.ResponseWriter, err error) {
	ce, ok := err.(*CustomError)
	if !ok {
		ce = &CustomError{Code: ErrInternalServerErrorCode, Message: "Internal server error"}
	}

	message, ok := errorMap[ce.Code]
	if !ok {
		message = ce.Message
	}

	// Set the appropriate HTTP status code based on the error code
	statusCode := http.StatusInternalServerError
	switch ce.Code {
	case ErrInvalidRequestCode:
		statusCode = http.StatusBadRequest
	case ErrResourceNotFoundCode:
		statusCode = http.StatusNotFound
	}

	w.WriteHeader(statusCode)
	fmt.Fprintf(w, "Error: %s\n", message)
}
