package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

// Custom errors for downstream service
var (
	ErrInvalidData        = errors.New("invalid data provided")
	ErrUserAlreadyExists = errors.New("user already exists")
)

// apiError represents a custom error response from the API
type apiError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// Error implements the error interface
func (e apiError) Error() string {
	return fmt.Sprintf("status=%d, message=%s", e.Status, e.Message)
}

// CallDownstreamService simulates calling a downstream service
func CallDownstreamService(user User) error {
	// Simulate intermediate failures
	if user.Username == "alice" {
		// Return a custom error with annotations
		return UserCreationError{Err: ErrUserAlreadyExists, Stage: "downstream", Description: "User already exists in the downstream service"}
	}
	if user.Username == "" {
		return UserCreationError{Err: ErrInvalidData, Stage: "downstream", Description: "Invalid data provided to the downstream service"}
	}
	return nil
}

// User represents a user to be created in the downstream service
type User struct {
	Username string `json:"username"`
}

// handleCreateUser handles the HTTP request for creating a user
func handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := CallDownstreamService(user)
	if err != nil {
		switch err.(type) {
		case UserCreationError:
			uce := err.(UserCreationError)
			handleDownstreamError(w, uce)
		default:
			handleDownstreamError(w, apiError{Status: http.StatusInternalServerError, Message: err.Error()})
		}
		return
	}

	// Success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, `{"message": "User created successfully"}`)
}

func handleDownstreamError(w http.ResponseWriter, err error) {
	switch err := err.(type) {
	case UserCreationError:
		// Map specific downstream errors to custom HTTP status codes
		statusCode := http.StatusInternalServerError
		if err.Err == ErrInvalidData {
			statusCode = http.StatusBadRequest
		} else if err.Err == ErrUserAlreadyExists {
			statusCode = http.StatusConflict
		}
		respondWithError(w, statusCode, err.Description)
	case apiError:
		respondWithError(w, err.Status, err.Message)
	default:
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
}

func respondWithError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(apiError{Status: statusCode, Message: message})
	if err != nil {
		log.Printf("Error encoding error response: %v", err)
	}
}

func main() {
	http.HandleFunc("/createUser", handleCreateUser)

	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
