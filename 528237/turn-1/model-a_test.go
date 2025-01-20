package main

import (
	"strconv"
	"testing"
)

// User struct with typical user details.
type User struct {
	Username string
	Email    string
	Age      int
}

// generateUsers generates a slice of User structs with dummy data.
func generateUsers(num int) []User {
	users := make([]User, num)
	for i := 0; i < num; i++ {
		users[i] = User{
			Username: "User" + strconv.Itoa(i),
			Email:    "user" + strconv.Itoa(i) + "@example.com",
			Age:      i % 100, // Example: age between 0 and 99.
		}
	}
	return users
}

// BenchmarkInsertAppend benchmarks inserting using append.
func BenchmarkInsertAppend(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var users []User
		for i := 0; i < 10000; i++ {
			users = append(users, User{
				Username: "User" + strconv.Itoa(i),
				Email:    "user" + strconv.Itoa(i) + "@example.com",
				Age:      i % 100,
			})
		}
	}
}

// BenchmarkInsertPreallocate benchmarks inserting with preallocated slice.
func BenchmarkInsertPreallocate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		users := make([]User, 0, 10000)
		for i := 0; i < 10000; i++ {
			users = append(users, User{
				Username: "User" + strconv.Itoa(i),
				Email:    "user" + strconv.Itoa(i) + "@example.com",
				Age:      i % 100,
			})
		}
	}
}

// BenchmarkInsertFixed benchmarks inserting with fixed-sized slice.
func BenchmarkInsertFixed(b *testing.B) {
	for n := 0; n < b.N; n++ {
		users := make([]User, 10000)
		for i := 0; i < 10000; i++ {
			users[i] = User{
				Username: "User" + strconv.Itoa(i),
				Email:    "user" + strconv.Itoa(i) + "@example.com",
				Age:      i % 100,
			}
		}
	}
}

func TestMain(m *testing.M) {
	m.Run()
}
