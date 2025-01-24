package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

// User holds user information
type User struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}

// Order holds order information
type Order struct {
	ID          int64   `json:"id"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
}

// KeyValueStore simulates a database with separate tables
type KeyValueStore struct {
	mu         *sync.RWMutex
	users      map[int64]User
	orders     map[int64]Order
	userOrders map[int64][]int64 // maps user IDs to order IDs
	files      map[string]string
}

// NewKeyValueStore initializes the key-value store
func NewKeyValueStore(userFile, orderFile, userOrderFile string) *KeyValueStore {
	kv := &KeyValueStore{
		users:      make(map[int64]User),
		orders:     make(map[int64]Order),
		userOrders: make(map[int64][]int64),
		files: map[string]string{
			"users":      userFile,
			"orders":     orderFile,
			"userOrders": userOrderFile,
		},
		mu: &sync.RWMutex{},
	}
	kv.load()
	return kv
}

// load loads data from files into the appropriate maps
func (kv *KeyValueStore) load() error {
	for key, file := range kv.files {
		data, err := os.ReadFile(file)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return err
		}
		switch key {
		case "users":
			json.Unmarshal(data, &kv.users)
		case "orders":
			json.Unmarshal(data, &kv.orders)
		case "userOrders":
			json.Unmarshal(data, &kv.userOrders)
		}
	}
	return nil
}

// save persists all data back to files
func (kv *KeyValueStore) save() error {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	for key, file := range kv.files {
		var data []byte
		var err error
		switch key {
		case "users":
			data, err = json.MarshalIndent(kv.users, "", "  ")
		case "orders":
			data, err = json.MarshalIndent(kv.orders, "", "  ")
		case "userOrders":
			data, err = json.MarshalIndent(kv.userOrders, "", "  ")
		}
		if err != nil {
			return err
		}
		if err := os.WriteFile(file, data, 0644); err != nil {
			return err
		}
	}
	return nil
}

// CRUD operations for Users
func (kv *KeyValueStore) CreateUser(user User) {
	kv.mu.Lock()
	kv.users[user.ID] = user
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
	delete(kv.users, id)
	delete(kv.userOrders, id)
	kv.mu.Unlock()
}

// CRUD operations for Orders
func (kv *KeyValueStore) CreateOrder(order Order) {
	kv.mu.Lock()
	kv.orders[order.ID] = order
	kv.mu.Unlock()
}

func (kv *KeyValueStore) GetOrder(id int64) (Order, bool) {
	kv.mu.RLock()
	defer kv.mu.RUnlock()
	order, ok := kv.orders[id]
	return order, ok
}

func (kv *KeyValueStore) DeleteOrder(id int64) {
	kv.mu.Lock()
	delete(kv.orders, id)
	// Remove order from all users
	for userID, orderIDs := range kv.userOrders {
		for i, oid := range orderIDs {
			if oid == id {
				kv.userOrders[userID] = append(orderIDs[:i], orderIDs[i+1:]...)
				break
			}
		}
	}
	kv.mu.Unlock()
}

// LinkUserToOrder establishes a many-to-many relationship
func (kv *KeyValueStore) LinkUserToOrder(userID, orderID int64) {
	kv.mu.Lock()
	if value, ok := kv.userOrders[userID]; ok {
		kv.userOrders[userID] = append(value, orderID)
	} else {
		kv.userOrders[userID] = []int64{orderID}
	}
	kv.mu.Unlock()
}

// Helper functions for generating random data
const alphabets = "abcdefghijklmnopqrstuvwxyz"

func randomString(size int) string {
	if size == 0 {
		return ""
	}

	buf := &bytes.Buffer{}
	defer buf.Reset()

	for size > 0 {
		buf.WriteByte(alphabets[rand.Intn(len(alphabets))])
		size--
	}
	return buf.String()
}

var emailHosts = [3]string{"example.com", "learning.tech", "test.io"}

func randomEmail(size int) string {
	if size == 0 {
		return ""
	}

	buf := &bytes.Buffer{}
	defer buf.Reset()

	for size > 0 {
		buf.WriteByte(alphabets[rand.Intn(len(alphabets))])
		size--
	}

	buf.WriteByte('@')
	buf.WriteString(emailHosts[rand.Intn(len(emailHosts))])

	return buf.String()
}

func generateRandomUsers() <-chan User {
	userInfoChan := make(chan User)
	// generating random records for 2 seconds
	ticker := time.NewTicker(2 * time.Second)

	go func() {
		defer close(userInfoChan)

		for {
			select {
			case <-ticker.C:
				return
			default:
				userInfoChan <- User{ID: rand.Int63n(1000), Email: randomEmail(rand.Intn(30)), Name: randomString(rand.Intn(30)), Age: rand.Intn(80)}
			}
		}
	}()

	return userInfoChan
}

func generateRandomOrders() <-chan Order {
	orderInfoChan := make(chan Order)
	// generating random records for 2 seconds
	ticker := time.NewTicker(2 * time.Second)

	go func() {
		defer close(orderInfoChan)

		for {
			select {
			case <-ticker.C:
				return
			default:
				orderInfoChan <- Order{ID: rand.Int63n(1000), Description: randomString(rand.Intn(80)), Amount: rand.Float64() * 100}
			}
		}
	}()

	return orderInfoChan
}

// Main function for demonstration
func main() {
	kvStore := NewKeyValueStore("users.json", "orders.json", "userOrders.json")
	defer kvStore.save()

	userInfoChan := generateRandomUsers()
	orderInfoChan := generateRandomOrders()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for userInfo := range userInfoChan {
			kvStore.CreateUser(userInfo)
		}
	}()
	go func() {
		defer wg.Done()
		for orderInfo := range orderInfoChan {
			kvStore.CreateOrder(orderInfo)
		}
	}()

	wg.Wait()

	// mapping user-orders
	size := 1000
	for size > 0 {
		randomUserID := rand.Int63n(1000)
		randomOrderID := rand.Int63n(1000)
		_, isUserExists := kvStore.GetUser(randomUserID)
		_, isOrderExists := kvStore.GetOrder(randomOrderID)

		if isUserExists && isOrderExists {
			kvStore.LinkUserToOrder(randomUserID, randomOrderID)
		}

		size--
	}

	fmt.Println("records created: ", len(kvStore.users), len(kvStore.orders), len(kvStore.userOrders))
}
