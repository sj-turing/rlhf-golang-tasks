package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// User represents a user in our table.
type User struct {
	ID   int
	Name string
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

func main() {
	// Create a new user table
	table := NewUserTable()

	// Simulate high insertions and deletions
	const numOperations = 100000
	const numUsers = 1000

	start := time.Now()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 1; i <= numOperations; i++ {
			table.Insert(User{ID: i, Name: fmt.Sprintf("User %d", i)})
		}
	}()

	go func() {
		defer wg.Done()
		for i := 1; i <= numOperations; i++ {
			if i%2 == 0 {
				table.Delete(i)
			}
		}
	}()

	wg.Wait()

	elapsed := time.Since(start)
	fmt.Printf("Elapsed time for %d operations: %s\n", numOperations, elapsed)

	// Analyze memory usage
	runtime.GC() // Force a garbage collection
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("Total Alloc: %d bytes\n", m.TotalAlloc)
	fmt.Printf("Sys: %d bytes\n", m.Sys)
	fmt.Printf("NumGC: %d\n", m.NumGC)
}
