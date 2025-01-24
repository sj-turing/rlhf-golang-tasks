package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

// User represents the user information to be stored
type User struct {
	Email string
	Name  string
	Age   int
}

// KeyValueStore handles the in-memory storage and file I/O for persistence
type KeyValueStore struct {
	data map[string]User
	mu   sync.RWMutex
	file string
}

// NewKeyValueStore initializes a new KeyValueStore
func NewKeyValueStore(filename string) *KeyValueStore {
	store := &KeyValueStore{
		data: make(map[string]User),
		file: filename,
	}
	store.loadFromFile()
	return store
}

// Set adds or updates a user with the given key
func (store *KeyValueStore) Set(key string, user User) {
	store.mu.Lock()
	defer store.mu.Unlock()
	store.data[key] = user
	store.saveToFile()
}

// Get retrieves a user by key if it exists
func (store *KeyValueStore) Get(key string) (User, bool) {
	store.mu.RLock()
	defer store.mu.RUnlock()
	user, exists := store.data[key]
	return user, exists
}

// Delete removes a user by key
func (store *KeyValueStore) Delete(key string) {
	store.mu.Lock()
	defer store.mu.Unlock()
	delete(store.data, key)
	store.saveToFile()
}

// loadFromFile reads data from the file and populates the map
func (store *KeyValueStore) loadFromFile() {
	file, err := ioutil.ReadFile(store.file)
	if err != nil {
		if os.IsNotExist(err) {
			return // File does not exist, nothing to load
		}
		fmt.Println("Error reading file:", err)
		return
	}

	err = json.Unmarshal(file, &store.data)
	if err != nil {
		fmt.Println("Error unmarshalling data:", err)
	}
}

// saveToFile writes data to the file for persistence
func (store *KeyValueStore) saveToFile() {
	store.mu.RLock()
	defer store.mu.RUnlock()
	data, err := json.Marshal(store.data)
	if err != nil {
		fmt.Println("Error marshalling data:", err)
		return
	}

	err = ioutil.WriteFile(store.file, data, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
	}
}

func main() {
	// Example usage
	store := NewKeyValueStore("store.json")

	// Adding some users
	store.Set("user1", User{Email: "user1@example.com", Name: "John Doe", Age: 30})
	store.Set("user2", User{Email: "user2@example.com", Name: "Jane Smith", Age: 25})

	// Retrieving a user
	user, exists := store.Get("user1")
	if exists {
		fmt.Printf("Retrieved: %+v\n", user)
	} else {
		fmt.Println("User not found")
	}

	// Deleting a user
	store.Delete("user2")

	// Attempt to retrieve deleted user
	user, exists = store.Get("user2")
	if exists {
		fmt.Printf("Retrieved: %+v\n", user)
	} else {
		fmt.Println("User not found")
	}
}
