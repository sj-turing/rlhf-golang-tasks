package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

// User represents a user in our database.
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// UserTable simulates a database table using a slice of User structs.
type UserTable struct {
	users []User
	mu    sync.Mutex // For thread safety
}

// NewUserTable creates a new UserTable instance.
func NewUserTable() *UserTable {
	return &UserTable{users: make([]User, 0)}
}

// Insert inserts a new user into the table.
func (ut *UserTable) Insert(user User) {
	ut.mu.Lock()
	defer ut.mu.Unlock()
	ut.users = append(ut.users, user)
}

// Delete deletes a user from the table by ID.
func (ut *UserTable) Delete(id int) {
	ut.mu.Lock()
	defer ut.mu.Unlock()
	for i, user := range ut.users {
		if user.ID == id {
			ut.users = append(ut.users[:i], ut.users[i+1:]...)
			return
		}
	}
}

// FindByID finds a user by ID.
func (ut *UserTable) FindByID(id int) *User {
	for _, user := range ut.users {
		if user.ID == id {
			return &user
		}
	}
	return nil
}

// All returns all users in the table.
func (ut *UserTable) All() []User {
	ut.mu.Lock()
	defer ut.mu.Unlock()
	return ut.users
}

// HandleGetUsers handles GET requests to retrieve all users.
func (ut *UserTable) HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	users := ut.All()
	json.NewEncoder(w).Encode(users)
}

// HandleGetUser handles GET requests to retrieve a single user by ID.
func (ut *UserTable) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	vars := map[string]string{}
	http.ParseQuery(r.URL.RawQuery)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "id parameter required", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id must be an integer", http.StatusBadRequest)
		return
	}

	user := ut.FindByID(id)
	if user == nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// HandlePostUser handles POST requests to insert a new user.
func (ut *UserTable) HandlePostUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "failed to parse JSON", http.StatusBadRequest)
		return
	}

	ut.Insert(user)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// HandleDeleteUser handles DELETE requests to remove a user by ID.
func (ut *UserTable) HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := map[string]string{}
	http.ParseQuery(r.URL.RawQuery)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "id parameter required", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id must be an integer", http.StatusBadRequest)
		return
	}

	ut.Delete(id)
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	table := NewUserTable()

	http.HandleFunc("/users", table.HandleGetUsers)
	http.HandleFunc("/users", table.HandlePostUser)
	http.HandleFunc("/users/", table.HandleGetUser)
	http.HandleFunc("/users/", table.HandleDeleteUser)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
