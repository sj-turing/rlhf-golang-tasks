package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

type User struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}

type Order struct {
	ID        int64  `json:"id"`
	UserID    int64  `json:"user_id"`
	OrderDate time.Time `json:"order_date"`
	Items     []string `json:"items"`
}

type UserOrders struct {
	UserID  int64 `json:"user_id"`
	OrderID int64 `json:"order_id"`
}

type KeyValueStore struct {
	mu        *sync.RWMutex
	users     map[int64]User
	orders    map[int64]Order
	userOrders map[int64][]int64
	files     []string
}

func NewKeyValueStore(usersFile, ordersFile, userOrdersFile string) *KeyValueStore {
	kv := &KeyValueStore{
		users:     make(map[int64]User),
		orders:    make(map[int64]Order),
		userOrders: make(map[int64][]int64),
		files:     []string{usersFile, ordersFile, userOrdersFile},
		mu:        &sync.RWMutex{},
	}

	for _, file := range kv.files {
		if err := kv.load(file); err != nil {
			log.Fatalf("Failed to load %s: %v", file, err)
		}
	}

	return kv
}
func (kv *KeyValueStore) load(file string) error {
	var data interface{}
	switch file {
	case kv.files[0]:
		data = &kv.users
	case kv.files[1]:
		data = &kv.orders
	case kv.files[2]:
		data = &kv.userOrders
	default:
		return fmt.Errorf("unknown file type: %s", file)
	}

	d, err := os.ReadFile(file)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	if err := json.Unmarshal(d, data); err != nil {
		return err
	}

	return nil
}

func (kv *KeyValueStore) save() error {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	for _, file := range kv.files {
		var data interface{}
		switch file {
		case kv.files[0]:
			data = kv.users
		case kv.files[1]:
			data = kv.orders
		case kv.files[2]:
			data = kv.userOrders
		}

		d, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			return err
		}

		if err := os.WriteFile(file, d, 0644); err != nil {
			return err
		}
	}

	return nil
}

func (kv *KeyValueStore) PutUser(id int64, user User) {
	kv.mu.Lock()
	kv.users[id] = user
	kv.mu.Unlock()
}

func (kv *KeyValueStore) GetUser(id int64) (User, bool) {
	kv.mu.RLock()
	defer kv.mu.RUnlock()
	user, ok := kv.users[id]
	return user, ok
}

func (kv *KeyValueStore) DeleteUser(id int64) {
	kv.mu.Lock()
	defer kv.mu.Unlock()
