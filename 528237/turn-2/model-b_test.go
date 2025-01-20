package main

import (
	"sync"
	"testing"
	"time"
)

// User struct with typical user details.
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

// GetUserDetails retrieves recently visited users within the last 12 hours.
func GetUserDetails(numUsers int) []*User {
	users := make([]*User, 0, numUsers)

	for i := 0; i < numUsers; i++ {
		user := userPool.Get().(*User)
		user.Username = "User" + strconv.Itoa(i)
		user.Email = "user" + strconv.Itoa(i) + "@example.com"
		user.Age = i % 100
		users = append(users, user)
	}

	// After processing, put the users back into the pool for reuse
	for _, user := range users {
		userPool.Put(user)
	}

	return users
}

// BenchmarkGetUserDetails benchmarks retrieving user details.
func BenchmarkGetUserDetails(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetUserDetails(10000)
	}
}
