package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Custom error types with annotations
type InvalidInputError struct {
	Err  error  // The underlying error
	Field string // The field that caused the error
}

func (e InvalidInputError) Error() string {
	return fmt.Sprintf("invalid input: field=%s, error=%v", e.Field, e.Err)
}

type UserAlreadyExistsError struct {
	Err        error  // The underlying error
	UserID     string // The ID of the existing user
	Username   string // The username of the existing user
}

func (e UserAlreadyExistsError) Error() string {
	return fmt.Sprintf("user already exists: userID=%s, username=%s, error=%v", e.UserID, e.Username, e.Err)
}

// User struct
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// SuccessResponse struct for custom successful response
type SuccessResponse struct {
	Message string `json:"message"`
	Data    interface{} `json:"data"`
}

// ErrorResponse struct for custom error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// Function to create a user
func createUser(username string, email string) (User, error) {
	// Simulate intermediate failures
	if username == "" {
		return User{}, InvalidInputError{Err: errors.New("username is required"), Field: "username"}
	}

	if email == "" {
		return User{}, InvalidInputError{Err: errors.New("email is required"), Field: "email"}
	}

	// Check if user already exists (simulated)
	if username == "existinguser" {
		return User{}, UserAlreadyExistsError{Err: errors.New("user already exists"), UserID: "1", Username: "existinguser"}
	}

	// Create the user (simulated)
	user := User{
		ID:       "1",
		Username: username,
		Email:    email,
	}
	return user, nil
}

// HTTP handler for user creation
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	// Decode the request body
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Create the user
	createdUser, err := createUser(user.Username, user.Email)
	if err != nil {
		switch err.(type) {
		case InvalidInputError:
			respondWithError(w, http.StatusBadRequest, err.Error())
		case UserAlreadyExistsError:
			respondWithError(w, http.StatusConflict, err.Error())
		default:
			respondWithError(w, http.StatusInternalServerError, "Internal server error")
		}
		return
	}

	// Respond with success
	respondWithSuccess(w, http.StatusCreated, "User created successfully", createdUser)
}

// Helper function to respond with a custom successful response
func respondWithSuccess(w http.ResponseWriter, status int, message string, data interface{}) {
	response := SuccessResponse{
		Message: message,
		Data:    data,
	}
	json.NewEncoder(w).Encode(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
}

// Helper function to respond with a custom error response
func respondWithError(w http.ResponseWriter, status int, message string) {
	response := ErrorResponse{
		Error:   http.StatusText(status),
