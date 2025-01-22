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
	mu    sync.RWMutex
	users map[string]*User
}

// NewUserService initializes and returns a new UserService
func NewUserService() *UserService {
	return &UserService{
		users: make(map[string]*User),
	}
}

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

	// Check if user exists
	s.mu.RLock()
	_, exists := s.users[user.Username]
	s.mu.RUnlock()

	if exists {
		respondWithJSON(w, http.StatusConflict, APIResponse{Status: "error", Message: "User already exists"})
		return
	}

	s.mu.Lock()
	s.users[user.Username] = &user
	s.mu.Unlock()

	respondWithJSON(w, http.StatusCreated, APIResponse{Status: "success", Message: "User created"})
}

// SignIn authenticates a user
func (s *UserService) SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithJSON(w, http.StatusMethodNotAllowed, APIResponse{Status: "error", Message: "Method not allowed"})
		return
	}
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		respondWithJSON(w, http.StatusBadRequest, APIResponse{Status: "error", Message: "Invalid request body"})
		return
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	user, ok := s.users[creds.Username]
	if !ok {
		respondWithJSON(w, http.StatusBadRequest, APIResponse{Status: "error", Message: "Invalid Username/Password"})
		return
	}

	if user.Password != creds.Password {
		respondWithJSON(w, http.StatusBadRequest, APIResponse{Status: "error", Message: "Invalid Username/Password"})
		return
	}

	respondWithJSON(w, http.StatusOK, APIResponse{Status: "success", Message: "User Login Response", Data: user})
}

// GetUserDetails gets details of a specified user
func (s *UserService) GetUserDetails(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	fmt.Println(username)
	s.mu.RLock()
	user, exists := s.users[username]
	s.mu.RUnlock()

	if !exists {
		respondWithJSON(w, http.StatusNotFound, APIResponse{Status: "error", Message: "User not found"})
		return
	}
	respondWithJSON(w, http.StatusOK, APIResponse{Status: "success", Message: "User Record", Data: user})
}

// UpdateUserAddress updates a user's address
func (s *UserService) UpdateUserAddress(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		respondWithJSON(w, http.StatusMethodNotAllowed, APIResponse{Status: "error", Message: "Method not allowed"})
		return
	}
	var update struct {
		Username string `json:"username"`
		Address  string `json:"address"`
	}
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		respondWithJSON(w, http.StatusBadRequest, APIResponse{Status: "error", Message: err.Error()})
		return
	}

	s.mu.Lock()
	user, exists := s.users[update.Username]
	if exists {
		user.Address = update.Address
	}
	s.mu.Unlock()

	if !exists {
		respondWithJSON(w, http.StatusNotFound, APIResponse{Status: "error", Message: "User not found"})
		return
	}
	respondWithJSON(w, http.StatusOK, APIResponse{Status: "success", Message: "User details successfully udpated", Data: user})
}

// GetUsersByAddress retrieves users based on the given address
func (s *UserService) GetUsersByAddress(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	if address == "" {
		respondWithJSON(w, http.StatusBadRequest, APIResponse{Status: "error", Message: "Address query param is required"})
		return
	}

	s.mu.RLock()
	defer s.mu.RUnlock()
	var users []*User
	for _, user := range s.users {
		if user.Address == address {
			users = append(users, user)
		}
	}
	if len(users) == 0 {
		respondWithJSON(w, http.StatusNotFound, APIResponse{Status: "error", Message: "No users found with the provided address"})
		return
	}

	respondWithJSON(w, http.StatusOK, APIResponse{Status: "success", Message: "Users List based on Address", Data: users})
}

func main() {
	userService := NewUserService()

	http.HandleFunc("/signup", userService.SignUp)
	http.HandleFunc("/signin", userService.SignIn)
	http.HandleFunc("/getuserdetails", userService.GetUserDetails)
	http.HandleFunc("/updateuseraddress", userService.UpdateUserAddress)
	http.HandleFunc("/getusersbyaddress", userService.GetUsersByAddress) // new endpoint

	log.Println("Server starting on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
