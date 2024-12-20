package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

// User represents a user
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SuccessResponse is a custom successful response
type SuccessResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	User    User   `json:"user"`
}

// ErrorResponse is a custom error response
type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// UserCreationError is a custom error type with annotations
type UserCreationError struct {
	Err         error  // The underlying error
	Stage       string // The stage of creation that failed
	Description string // Additional description of the failure
}

// Error implements the error interface
func (e UserCreationError) Error() string {
	return fmt.Sprintf("user creation error: stage=%s, description=%s, error=%v", e.Stage, e.Description, e.Err)
}

// CreateUser is a function that simulates user creation with annotations
func CreateUser(username, email, password string) error {
	// Simulate validation errors
	if username == "" {
		return UserCreationError{Err: errors.New("username is required"), Stage: "validation", Description: "Username is missing"}
	}
	if email == "" {
		return UserCreationError{Err: errors.New("email is required"), Stage: "validation", Description: "Email is missing"}
	}

	// Simulate a database error during user creation
	userID := "123"
	user := User{ID: userID, Username: username, Email: email, Password: password}
	if _, err := saveUserToDatabase(user); err != nil {
		return UserCreationError{Err: err, Stage: "database", Description: "Failed to save user to database"}
	}

	return nil
}

func saveUserToDatabase(user User) (string, error) {
	// Simulate a database operation
	return user.ID, nil
}

// handleCreateUser is the HTTP handler for user creation
func handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := CreateUser(user.Username, user.Email, user.Password)
	if err != nil {
		// Handle user creation error
		switch err.(type) {
		case UserCreationError:
			uce := err.(UserCreationError)
			http.Error(w, json.Marshal(ErrorResponse{Status: "error", Message: uce.Description}), http.StatusInternalServerError)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Handle successful user creation
	successResponse := SuccessResponse{Status: "success", Message: "User created successfully", User: user}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(successResponse); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/createUser", handleCreateUser)

	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
