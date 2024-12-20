package main

import (
	"errors"
	"fmt"
	"log"
)

// UserSessionError represents a custom error type for user session-related issues
type UserSessionError struct {
	Error   error
	Code    int
	Message string
	Details string
}

// Error implements the error interface
func (e *UserSessionError) Error() string {
	return fmt.Sprintf("UserSessionError[code=%d, message=%s, details=%s]: %v",
		e.Code, e.Message, e.Details, e.Error)
}

// CreateSession creates a new user session
func CreateSession(userID string) (*UserSessionError, error) {
	if userID == "" {
		return &UserSessionError{
			Error:   errors.New("invalid user ID"),
			Code:    10001,
			Message: "User ID is required",
			Details: "user ID must not be empty",
		}, nil
	}

	// Simulating an error during session creation
	return nil, errors.New("failed to create session")
}

func main() {
	userSessionError, err := CreateSession("")
	if err != nil {
		if userSessionError != nil {
			// Log the user session error with details
			log.Print(userSessionError)
		} else {
			// Log a generic error
			log.Printf("An unexpected error occurred: %v", err)
		}
	} else {
		log.Print("Session created successfully")
	}
}
