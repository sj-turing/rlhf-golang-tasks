// filename: main.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

// UserService to handle operations on user data
type UserService struct {
	mu             sync.RWMutex
	users          map[string]*User
	usersByAddress map[string][]*User
}

// NewUserService initializes and returns a new UserService
func NewUserService() *UserService {
	return &UserService{
		users:          make(map[string]*User),
		usersByAddress: make(map[string][]*User),
	}
}

// Adds a new user to the user service and updates the usersByAddress index
func (s *UserService) AddUser(user *User) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.users[user.Username] = user
	s.usersByAddress[user.Address] = append(s.usersByAddress[user.Address], user)
}

// Removes a user from the user service and updates the usersByAddress index
func (s *UserService) RemoveUser(username string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	user, ok := s.users[username]
	if !ok {
		return
	}
	delete(s.users, username)
	for i, u := range s.usersByAddress[user.Address] {
		if u.Username == username {
			s.usersByAddress[user.Address] = append(s.usersByAddress[user.Address][:i], s.usersByAddress[user.Address][i+1:]...)
			return
		}
	}
}

// Gets a user by username from the user service
func (s *UserService) GetUser(username string) *User {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.users[username]
}

// Gets users by address from the usersByAddress index
func (s *UserService) GetUsersByAddress(address string) []*User {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.usersByAddress[address]
}

// APIResponse struct which wraps the API response
type APIResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// respondWithJSON returns response with content-type JSON
func respondWithJSON(w http.ResponseWriter, status int, resp APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
}

// SignUp creates a new user
func (s *UserService) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithJSON(w, http.StatusMethodNotAllowed, APIResponse{Status: "error", Message: "Method not allowed"})
		return
	}
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithJSON(w, http.StatusBadRequest, APIResponse{Status: "error", Message: "Invalid request body"})
		return
	}
	s.AddUser(&user)
	respondWithJSON(w, http.StatusCreated, APIResponse{Status: "success", Message: "User created"})
}

// SignIn authenticates a user
func (s *UserService) SignIn(w http.ResponseWriter, r *http.Request) {
	// Existing SignIn logic remains the same
}

// GetUserDetails gets details of a specified user
func (s *UserService) GetUserDetails(w http.ResponseWriter, r *http.Request) {
	// Existing GetUserDetails logic remains the same
}
