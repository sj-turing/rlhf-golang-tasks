package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
	"time"

	_ "net/http/pprof"
)

type User struct {
	Username string
	Email    string
	Age      int
}

var (
	userPool = sync.Pool{
		New: func() interface{} {
			return new(User)
		},
	}
)

func GetUserDetails(w http.ResponseWriter, r *http.Request) {
	numUsers, err := strconv.Atoi(r.URL.Query().Get("num"))
	if err != nil || numUsers <= 0 {
		http.Error(w, "Invalid number of users", http.StatusBadRequest)
		return
	}

	users := make([]*User, 0, numUsers)
	for i := 0; i < numUsers; i++ {
		user := userPool.Get().(*User)
		user.Username = "User" + strconv.Itoa(i)
		user.Email = "user" + strconv.Itoa(i) + "@example.com"
		user.Age = i % 100
		users = append(users, user)
	}

	response, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)

	// Return users to pool
	for _, user := range users {
		userPool.Put(user)
	}
}

func main() {
	http.HandleFunc("/userdetails", GetUserDetails)
	// Register pprof handlers
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	http.ListenAndServe(":8080", nil)
}
