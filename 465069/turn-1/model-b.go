package main

import (
	"errors"
	"fmt"
)

// SessionError is a custom error type with annotations
type SessionError struct {
	Err       error  // The underlying error
	Action    string // The action that caused the error
	SessionID string // The ID of the user session
}

// Error implements the error interface
func (e SessionError) Error() string {
	return fmt.Sprintf("session error: action=%s, sessionID=%s, error=%v", e.Action, e.SessionID, e.Err)
}

// Example function that uses annotations
func performUserAction(sessionID string, action string) error {
	// Simulate an error condition
	if action == "login" {
		return SessionError{Err: errors.New("invalid credentials"), Action: action, SessionID: sessionID}
	}
	return nil
}

func main() {
	sessionID := "abc123"
	action := "login"
	err := performUserAction(sessionID, action)
	if err != nil {
		switch err.(type) {
		case SessionError:
			// Handle session-specific errors
			se := err.(SessionError)
			fmt.Printf("Session Error: %v\n", se)
		default:
			// Handle other errors
			fmt.Printf("Unexpected Error: %v\n", err)
		}
	}
}
