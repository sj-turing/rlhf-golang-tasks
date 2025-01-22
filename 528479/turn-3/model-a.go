package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

type UserService struct {
	mu           sync.RWMutex
	users        map[string]*User
	addressIndex map[string][]*User
}

func NewUserService() *UserService {
	return &UserService{
		users:        make(map[string]*User),
		addressIndex: make(map[string][]*User),
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

	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.users[user.Username]; exists {
		respondWithJSON(w, http.StatusConflict, APIResponse{Status: "error", Message: "User already exists"})
		return
	}

	s.users[user.Username] = &user
	s.addressIndex[user.Address] = append(s.addressIndex[user.Address], &user)

	respondWithJSON(w, http.StatusCreated, APIResponse{Status: "success", Message: "User created"})
}

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
	if !ok || user.Password != creds.Password {
		respondWithJSON(w, http.StatusBadRequest, APIResponse{Status: "error", Message: "Invalid Username/Password"})
		return
	}

	respondWithJSON(w, http.StatusOK, APIResponse{Status: "success", Message: "User Login successful", Data: user})
}

func (s *UserService) GetUserDetails(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	s.mu.RLock()
	user, exists := s.users[username]
	s.mu.RUnlock()

	if !exists {
		respondWithJSON(w, http.StatusNotFound, APIResponse{Status: "error", Message: "User not found"})
		return
	}
	respondWithJSON(w, http.StatusOK, APIResponse{Status: "success", Message: "User record found", Data: user})
}

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
		respondWithJSON(w, http.StatusBadRequest, APIResponse{Status: "error", Message: "Invalid request body"})
		return
	}

	s.mu.Lock()
	user, exists := s.users[update.Username]
	if exists {
		// Remove user from old address index
		oldAddress := user.Address
		s.addressIndex[oldAddress] = removeUserFromSlice(s.addressIndex[oldAddress], user)

		// Update user address and update index
		user.Address = update.Address
		s.addressIndex[update.Address] = append(s.addressIndex[update.Address], user)
	}
	s.mu.Unlock()

	if !exists {
		respondWithJSON(w, http.StatusNotFound, APIResponse{Status: "error", Message: "User not found"})
		return
	}
	respondWithJSON(w, http.StatusOK, APIResponse{Status: "success", Message: "Address updated successfully", Data: user})
}

func removeUserFromSlice(users []*User, userToRemove *User) []*User {
	for i, user := range users {
		if user == userToRemove {
			return append(users[:i], users[i+1:]...)
		}
	}
	return users
}

func (s *UserService) GetUsersByAddress(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	if address == "" {
		respondWithJSON(w, http.StatusBadRequest, APIResponse{Status: "error", Message: "Address query param required"})
		return
	}

	s.mu.RLock()
	users, exists := s.addressIndex[address]
	s.mu.RUnlock()

	if !exists || len(users) == 0 {
		respondWithJSON(w, http.StatusNotFound, APIResponse{Status: "error", Message: "No users found with the provided address"})
		return
	}

	respondWithJSON(w, http.StatusOK, APIResponse{Status: "success", Message: "Users found", Data: users})
}

func main() {
	userService := NewUserService()

	http.HandleFunc("/signup", userService.SignUp)
	http.HandleFunc("/signin", userService.SignIn)
	http.HandleFunc("/getuserdetails", userService.GetUserDetails)
	http.HandleFunc("/updateuseraddress", userService.UpdateUserAddress)
	http.HandleFunc("/getusersbyaddress", userService.GetUsersByAddress)

	log.Println("Server starting on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
