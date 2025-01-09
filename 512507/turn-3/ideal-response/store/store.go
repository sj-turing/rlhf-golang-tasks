package store

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
	users    *sync.Map
	filename string
	mu       sync.Mutex
}

func NewUserStore(filename string) (*UserStore, error) {
	store := &UserStore{
		users:    new(sync.Map),
		filename: filename,
	}

	// Load existing data from file
	if err := store.load(); err != nil {
		return nil, err
	}

	gob.Register(User{})

	return store, nil
}

func (s *UserStore) save() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	file, err := os.OpenFile(s.filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	temp := map[string]User{}
	s.users.Range(func(k, v interface{}) bool {
		temp[k.(string)] = v.(User)
		return true
	})

	if err = gob.NewEncoder(file).Encode(temp); err != nil {
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

	var temp map[string]User
	if err := gob.NewDecoder(file).Decode(&temp); err != nil {
		if errors.Is(err, io.EOF) {
			return nil
		}
		return fmt.Errorf("failed to decode data: %v", err)
	}

	for key, value := range temp {
		s.users.Store(key, value)
	}

	return nil
}

func (s *UserStore) Create(email string, name string, age int) error {
	if _, ok := s.users.Load(email); ok {
		return errors.New("user with given email already exists")
	}

	s.users.Store(email, User{Name: name, Age: age})
	return s.save()
}

func (s *UserStore) Read(email string) (User, error) {
	user, ok := s.users.Load(email)
	if !ok {
		return User{}, errors.New("user not found")
	}

	return user.(User), nil
}

func (s *UserStore) Update(email string, name string, age int) error {
	if _, ok := s.users.Load(email); !ok {
		return errors.New("user not found")
	}

	s.users.Swap(email, User{Name: name, Age: age})

	return s.save()
}

func (s *UserStore) Delete(email string) error {
	if _, ok := s.users.Load(email); !ok {
		return errors.New("user not found")
	}

	s.users.Delete(email)

	return s.save()
}

func (s *UserStore) Save() error {
	fmt.Println("Saving user records into file")
	return s.save()
}
