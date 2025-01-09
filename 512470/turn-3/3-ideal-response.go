package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
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

type errorResponse struct {
	ErrorMessage string `json:"error_message"`
}

func respond(w http.ResponseWriter, code int, msg any) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(msg)
}

// HandleGetUsers handles GET requests to retrieve all users.
func (ut *UserTable) HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	users := ut.All()
	respond(w, http.StatusOK, users)
}

// HandleGetUser handles GET requests to retrieve a single user by ID.
func (ut *UserTable) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		respond(w, http.StatusBadRequest, errorResponse{"id must be an integer"})
		return
	}

	user := ut.FindByID(id)
	if user == nil {
		respond(w, http.StatusNotFound, errorResponse{"user not found"})
		return
	}

	respond(w, http.StatusOK, user)
}

// HandlePostUser handles POST requests to insert a new user.
func (ut *UserTable) HandlePostUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respond(w, http.StatusBadRequest, errorResponse{"failed to parse JSON"})
		return
	}

	ut.Insert(user)
	respond(w, http.StatusCreated, user)
}

// HandleDeleteUser handles DELETE requests to remove a user by ID.
func (ut *UserTable) HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		respond(w, http.StatusBadRequest, errorResponse{"id must be an integer"})
		return
	}

	user := ut.FindByID(id)
	if user == nil {
		respond(w, http.StatusNotFound, errorResponse{"user not found"})
		return
	}

	ut.Delete(id)
	respond(w, http.StatusOK, user)
}

func main() {
	router := mux.NewRouter()
	table := NewUserTable()

	router.HandleFunc("/user/{id}", table.HandleGetUser).Methods("GET")
	router.HandleFunc("/user", table.HandlePostUser).Methods("POST")
	router.HandleFunc("/users", table.HandleGetUsers).Methods("GET")
	router.HandleFunc("/user/{id}", table.HandleDeleteUser).Methods("DELETE")

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
