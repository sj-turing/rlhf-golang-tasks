package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

type User struct {
	Username string
	Password string
	Email    string
	Address  string
}

var (
	users      = make(map[string]User)
	usersMutex sync.RWMutex
)

type APIResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func respondWithJSON(w http.ResponseWriter, status int, resp APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithJSON(w, http.StatusMethodNotAllowed, APIResponse{Status: "error", Message: "Method not allowed"})
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithJSON(w, http.StatusBadRequest, APIResponse{Status: "error", Message: "Invalid request payload"})
		return
	}

	// Check if user exists
	usersMutex.RLock()
	_, exists := users[user.Username]
	usersMutex.RUnlock()

	if exists {
		respondWithJSON(w, http.StatusConflict, APIResponse{Status: "error", Message: "User already exists"})
		return
	}

	// Adding a new user
	usersMutex.Lock()
	users[user.Username] = user
	usersMutex.Unlock()

	respondWithJSON(w, http.StatusCreated, APIResponse{Status: "success", Message: "User created"})
}

func main() {
	http.HandleFunc("/signup", SignUpHandler)
	// Add handlers for signin, get user details, and update address

	log.Println("Server starting on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
