package main

import (
	"fmt"
	"testing"
)

type User struct {
	Username string
	Email    string
	Age      int
}

const numUsers = 1000000

func BenchmarkMemoryUsageWithSlice(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var users []User
		for i := 0; i < numUsers; i++ {
			user := User{
				Username: fmt.Sprintf("user%d", i),
				Email:    fmt.Sprintf("user%d@example.com", i),
				Age:      i % 100,
			}
			users = append(users, user)
		}
	}
}

func BenchmarkMemoryUsageWithMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		users := make(map[string]User, numUsers)
		for i := 0; i < numUsers; i++ {
			user := User{
				Username: fmt.Sprintf("user%d", i),
				Email:    fmt.Sprintf("user%d@example.com", i),
				Age:      i % 100,
			}
			users[user.Username] = user
		}
	}
}

func main() {
}
