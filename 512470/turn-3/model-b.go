package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"sync"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserTable struct {
	users map[int]*User
	mu    sync.RWMutex
}

func NewUserTable() *UserTable {
	return &UserTable{users: make(map[int]*User)}
}

func (ut *UserTable) Insert(user User) {
	ut.mu.Lock()
	defer ut.mu.Unlock()
	ut.users[user.ID] = &user
}

func (ut *UserTable) Delete(id int) {
	ut.mu.Lock()
	defer ut.mu.Unlock()
	delete(ut.users, id)
}

func (ut *UserTable) FindByID(id int) *User {
	ut.mu.RLock()
	defer ut.mu.RUnlock()
	return ut.users[id]
}

func (ut *UserTable) All() []User {
	ut.mu.RLock()
	defer ut.mu.RUnlock()

	users := make([]User, 0, len(ut.users))
	for _, user := range ut.users {
		users = append(users, *user)
	}
	return users
}

var (
	userTable = NewUserTable()
)

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	user := userTable.FindByID(id)
	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userTable.Insert(user)
	json.NewEncoder(w).Encode(user)
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	userTable.Delete(id)
	fmt.Fprintf(w, "User with ID %d deleted", id)
}

func listUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := userTable.All()
	json.NewEncoder(w).Encode(users)
}

func main() {
	http.HandleFunc("/get", getUserHandler)
	http.HandleFunc("/create", createUserHandler)
	http.HandleFunc("/delete", deleteUserHandler)
	http.HandleFunc("/list", listUsersHandler)

	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
