package store

import (
	"errors"
	"fmt"
	"sync"
)

type User struct {
	Name string
	Age  int
}

type UserStore struct {
	users    *sync.Map
	filename string
}

func NewUserStore(filename string) (*UserStore, error) {
	store := &UserStore{
		users:    new(sync.Map),
		filename: filename,
	}

	if err := store.load(); err != nil {
		return nil, err
	}

	return store, nil
}

func (s *UserStore) load() error {
	// Load existing data from file
}

func (s *UserStore) save() error {
	// Save current state of the map to the file
}

// Thread-safe CRUD operations
func (s *UserStore) Create(email string, name string, age int) error {
	// Implement create operation
}

func (s *UserStore) Read(email string) (User, error) {
	// Implement read operation
}

func (s *UserStore) Update(email string, name string, age int) error {
	// Implement update operation
}

func (s *UserStore) Delete(email string) error {
	// Implement delete operation
}
