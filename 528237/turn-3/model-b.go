package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

var users []User
var mu sync.RWMutex

func init() {
	users = make([]User, 0, 10000)
	// Populate the users slice with dummy data
	for i := 0; i < 10000; i++ {
		user := User{
			Username: fmt.Sprintf("User%d", i),
			Email:    fmt.Sprintf("user%d@example.com", i),
			Age:      i % 100,
		}
		users = append(users, user)
	}
}

func getUserDetailsHandler(w http.ResponseWriter, r *http.Request) {
	mu.RLock()
	defer mu.RUnlock()
	// Simulate some processing time
	time.Sleep(10 * time.Millisecond)
	// Return the list of users in JSON format
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(users)
}

func main() {
	http.HandleFunc("/users", getUserDetailsHandler)
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
