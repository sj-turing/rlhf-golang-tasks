package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type User struct {
	Id   int
	Name string
}

var (
	wg    sync.WaitGroup
	users sync.Map
)

func generateRandomUser() User {
	rand.Seed(time.Now().UnixNano())
	name := fmt.Sprintf("User %d", rand.Intn(1000))
	return User{Id: rand.Intn(10000), Name: name}
}
func addUser(user User) {
	defer wg.Done()
	users.Store(user.Id, user)
}
func getUser(id int) (User, bool) {
	defer wg.Done()
	user, ok := users.Load(id)
	return user.(User), ok
}
func deleteUser(id int) {
	defer wg.Done()
	users.Delete(id)
}
func main() {
	numReaders := 10
	numWriters := 5
	numOperations := 10000
	// Start worker goroutines
	wg.Add(numReaders + numWriters + numOperations)
	for i := 0; i < numWriters; i++ {
		go func() {
			for j := 0; j < numOperations/numWriters; j++ {
				user := generateRandomUser()
				addUser(user)
			}
		}()
	}
	for i := 0; i < numReaders; i++ {
		go func() {
			for j := 0; j < numOperations/numReaders; j++ {
				userId := rand.Intn(10000)
				user, ok := getUser(userId)
				if ok {
					fmt.Printf("Read User: %+v\n", user)
				} else {
					fmt.Printf("User with ID %d not found.\n", userId)
				}
			}
		}()
	}
	for i := 0; i < numOperations/2; i++ {
		go func() {
			userId := rand.Intn(10000)
			deleteUser(userId)
		}()
	}
	// Wait for all goroutines to complete
	wg.Wait()
}
