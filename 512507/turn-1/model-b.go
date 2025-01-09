package main

import (
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

type User struct {
	Name     string
	Age      int
	Location string
}

var users map[string]User
var mu sync.RWMutex

func init() {
	users = make(map[string]User)
	// Load the saved data from the file on application start
	if err := loadUsersFromFile(); err != nil {
		log.Printf("Error loading users: %v", err)
	}
}
func saveUsersToFile() error {
	mu.RLock()
	defer mu.RUnlock()
	f, err := os.Create("users.db")
	if err != nil {
		return err
	}
	defer f.Close()
	e := gob.NewEncoder(f)
	return e.Encode(users)
}
func loadUsersFromFile() error {
	f, err := os.Open("users.db")
	if os.IsNotExist(err) {
		return nil // File not found, it's fine
	}
	if err != nil {
		return err
	}
	defer f.Close()
	d := gob.NewDecoder(f)
	if err := d.Decode(&users); err != nil {
		return err
	}
	return nil
}
func StoreUser(email string, user User) error {
	mu.Lock()
	defer mu.Unlock()
	users[email] = user
	return saveUsersToFile()
}
func GetUser(email string) (User, error) {
	mu.RLock()
	defer mu.RUnlock()
	user, ok := users[email]
	if !ok {
		return User{}, fmt.Errorf("user with email %s not found", email)
	}
	return user, nil
}
func UpdateUser(email string, updates User) error {
	mu.Lock()
	defer mu.Unlock()
	if _, ok := users[email]; !ok {
		return fmt.Errorf("user with email %s not found", email)
	}
	// Update only the fields that were provided in the updates map
	if updates.Name != "" {
		users[email].Name = updates.Name
	}
	if updates.Age != 0 {
		users[email].Age = updates.Age
	}
	if updates.Location != "" {
		users[email].Location = updates.Location
	}
	return saveUsersToFile()
}
func DeleteUser(email string) error {
	mu.Lock()
	defer mu.Unlock()
	if _, ok := users[email]; !ok {
		return fmt.Errorf("user with email %s not found", email)
	}
	delete(users, email)
	return saveUsersToFile()
}
func main() {
	err := StoreUser("user1@example.com", User{Name: "User One", Age: 30, Location: "New York"})
	if err != nil {
		log.Fatalf("Error storing user: %v", err)
	}
	user, err := GetUser("user1@example.com")
	if err != nil {
		log.Fatalf("Error getting user: %v", err)
	}
	fmt.Println("User:", user)
	err = UpdateUser("user1@example.com", User{Location: "San Francisco"})
	if err != nil {
		log.Fatalf("Error updating user: %v", err)
	}
	err = DeleteUser("user1@example.com")
	if err != nil {
		log.Fatalf("Error deleting user: %v", err)
	}
}
