package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

// User structure
type User struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}

// Order structure
type Order struct {
	ID          int64   `json:"id"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
}

// KeyValueStore to hold users, orders, and userOrders
type KeyValueStore struct {
	mu         *sync.RWMutex
	users      map[int64]User
	orders     map[int64]Order
	userOrders map[int64][]int64
	files      map[string]string
}

// NewKeyValueStore initializes the store and loads data
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

// load data from files
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

// save data to files
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

// Create user
func (kv *KeyValueStore) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.ID = time.Now().UnixNano()
	kv.mu.Lock()
	kv.users[user.ID] = user
	kv.mu.Unlock()
	json.NewEncoder(w).Encode(map[string]int64{"userID": user.ID})
}

// Create order
func (kv *KeyValueStore) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	order.ID = time.Now().UnixNano()
	kv.mu.Lock()
	kv.orders[order.ID] = order
	kv.mu.Unlock()
	json.NewEncoder(w).Encode(map[string]int64{"orderID": order.ID})
}

// Link user with order
func (kv *KeyValueStore) LinkUserOrder(w http.ResponseWriter, r *http.Request) {
	var data struct {
		UserID  int64 `json:"userID"`
		OrderID int64 `json:"orderID"`
	}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	kv.mu.Lock()
	defer kv.mu.Unlock()

	if _, userExists := kv.users[data.UserID]; !userExists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	if _, orderExists := kv.orders[data.OrderID]; !orderExists {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}
	kv.userOrders[data.UserID] = append(kv.userOrders[data.UserID], data.OrderID)
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

// Retrieve all orders for a user
func (kv *KeyValueStore) GetUserOrders(w http.ResponseWriter, r *http.Request) {
	var data struct {
		UserID int64 `json:"userID"`
	}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	kv.mu.RLock()
	defer kv.mu.RUnlock()

	orderIDs, userExists := kv.userOrders[data.UserID]
	if !userExists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	orders := make([]Order, 0, len(orderIDs))
	for _, orderID := range orderIDs {
		if order, orderExists := kv.orders[orderID]; orderExists {
			orders = append(orders, order)
		}
	}
	json.NewEncoder(w).Encode(orders)
}

func main() {
	kvStore := NewKeyValueStore("users.json", "orders.json", "userOrders.json")
	defer kvStore.save()

	http.HandleFunc("/createUser", kvStore.CreateUser)
	http.HandleFunc("/createOrder", kvStore.CreateOrder)
	http.HandleFunc("/linkUserOrder", kvStore.LinkUserOrder)
	http.HandleFunc("/getUserOrders", kvStore.GetUserOrders)

	fmt.Println("Server is starting at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
