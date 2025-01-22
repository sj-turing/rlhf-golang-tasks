package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Address string `json:"address"`
}

// In-memory user store using a sync.Map for thread safety
var userStore sync.Map

// API handlers
func signUpHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	// Generate a unique ID for the user (using a simple UUID generator)
	user.ID = generateUUID()

	// Store the user in the in-memory store
	userStore.Store(user.ID, user)

	// Return the user data with the assigned ID
	json.NewEncoder(w).Encode(user)
}

func signInHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	// Perform authentication based on email and password (not implemented in this example)
	// For simplicity, we'll just check if the user exists in the store

	if _, ok := userStore.Load(user.ID); ok {
		// User found, return the user data
		json.NewEncoder(w).Encode(user)
	} else {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	}
}

func getUserDetailsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	// Load the user from the in-memory store using the provided ID
	user, ok := userStore.Load(id)
	if !ok {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user.(User))
}

func updateUserAddressHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	// Load the existing user from the store
	existingUser, ok := userStore.Load(user.ID)
	if !ok {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	// Update the address in the existing user data
	existingUser.(*User).Address = user.Address
	
	// Store the updated user back to the store
	userStore.Store(user.ID, existingUser)
	json.NewEncoder(w).Encode(existingUser.(User))
}

func main() {
	// Initialize the router
	router := http.NewServeMux()
	router.HandleFunc("/signup", signUpHandler)
	router.HandleFunc("/signin", signInHandler)
