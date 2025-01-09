package main

import (
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
)

type User struct {
	Name string
	Age  int
}

type UserStore struct {
	users    map[string]User
	mu       sync.RWMutex
	filename string
}

func NewUserStore(filename string) (*UserStore, error) {
	store := &UserStore{
		users:    make(map[string]User),
		mu:       sync.RWMutex{},
		filename: filename,
	}

	// Load existing data from file
	if err := store.load(); err != nil {
		return nil, err
	}

	gob.Register(&User{})

	return store, nil
}

func (s *UserStore) save() error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	file, err := os.OpenFile(s.filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(s.users)
	if err != nil {
		return fmt.Errorf("failed to encode data: %v", err)
	}

	return nil
}

func (s *UserStore) load() error {
	file, err := os.OpenFile(s.filename, os.O_RDONLY, 0644)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&s.users)
	if err != nil {
		return fmt.Errorf("failed to decode data: %v", err)
	}

	return nil
}

func (s *UserStore) Create(email string, name string, age int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.users[email]; ok {
		return errors.New("user with given email already exists")
	}
	s.users[email] = User{Name: name, Age: age}

	return s.save()
}

func (s *UserStore) Read(email string) (*User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	user, ok := s.users[email]
	if !ok {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (s *UserStore) Update(email string, name string, age int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.users[email]; !ok {
		return errors.New("user not found")
	}
	s.users[email] = User{Name: name, Age: age}

	return s.save()
}

func (s *UserStore) Delete(email string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.users[email]; !ok {
		return errors.New("user not found")
	}
	delete(s.users, email)

	return s.save()
}

func main() {
	filename := "users.gob"
	store, err := NewUserStore(filename)
	if err != nil {
		fmt.Println("Error creating user store:", err)
		return
	}

	// Example CRUD operations
	go func() {
		err := store.Create("charlie@example.com", "Charlie", 35)
		if err != nil {
			fmt.Println("Error creating user:", err)
		} else {
			fmt.Println("User created successfully")
		}
	}()

	go func() {
		user, err := store.Read("alice@example.com")
		if err != nil {
			fmt.Println("Error reading user:", err)
		} else {
			fmt.Println("Read user:", user)
		}
	}()

	go func() {
		err := store.Update("bob@example.com", "Bobby", 26)
		if err != nil {
			fmt.Println("Error updating user:", err)
		} else {
			fmt.Println("User updated successfully")
		}
	}()

	go func() {
		err := store.Delete("charlie@example.com")
		if err != nil {
			fmt.Println("Error deleting user:", err)
		} else {
			fmt.Println("User deleted successfully")
		}
	}()
}
