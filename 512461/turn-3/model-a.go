package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type User struct {
	ID   int
	Name string
}

var (
	wg             sync.WaitGroup
	users          []User
	mu             sync.RWMutex
	numGenerators  = 5
	numProcessors  = 3
	numUsersToLoad = 1000
)

func generateRandomUser() User {
	name := fmt.Sprintf("User_%03d", rand.Intn(1000))
	return User{ID: rand.Intn(10000), Name: name}
}

func generateUsersConcurrently() {
	defer wg.Done()
	for i := 0; i < numUsersToLoad; i++ {
		user := generateRandomUser()
		mu.Lock()
		users = append(users, user)
		mu.Unlock()
	}
}

func readUsersConcurrently() {
	defer wg.Done()
	for {
		time.Sleep(time.Millisecond * 100)
		mu.RLock()
		defer mu.RUnlock()
		if len(users) > 0 {
			fmt.Println("Current User List:")
			for _, user := range users {
				fmt.Printf("ID: %d, Name: %s\n", user.ID, user.Name)
			}
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	wg.Add(numGenerators + numProcessors + 1)

	// Start generating users concurrently
	for i := 0; i < numGenerators; i++ {
		go generateUsersConcurrently()
	}

	// Start reading users concurrently
	for i := 0; i < numProcessors; i++ {
		go readUsersConcurrently()
	}

	// Wait for all worker goroutines to complete
	wg.Wait()

	// At this point, all users have been generated
	fmt.Println("\nFinal User List:")
	mu.RLock()
	defer mu.RUnlock()
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s\n", user.ID, user.Name)
	}
}
