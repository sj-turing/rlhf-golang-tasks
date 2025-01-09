package main

import (
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"os"
)

type User struct {
	Name string
	Age  int
}

func saveUsersToFile(users map[string]User, filename string) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(users)
	if err != nil {
		return fmt.Errorf("failed to encode data: %v", err)
	}

	return nil
}

func loadUsersFromFile(filename string) (map[string]User, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	users := make(map[string]User)
	err = decoder.Decode(&users)
	if err != nil {
		return nil, fmt.Errorf("failed to decode data: %v", err)
	}

	return users, nil
}

func main() {
	// Create sample user data
	users := map[string]User{
		"alice@example.com": {Name: "Alice", Age: 30},
		"bob@example.com":   {Name: "Bob", Age: 25},
	}

	// Serialize user data to a file
	filename := "users.gob"
	err := saveUsersToFile(users, filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("User data saved to", filename)

	// Deserialize user data from the file
	loadedUsers, err := loadUsersFromFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("Loaded user data:", loadedUsers)
}
