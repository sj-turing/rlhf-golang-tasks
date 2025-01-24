package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

type User struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}

type Order struct {
	ID          int64   `json:"id"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
}

type KeyValueStore struct {
	mu         *sync.RWMutex
	users      map[int64]User
	orders     map[int64]Order
	userOrders map[int64][]int64
	files      map[string]string
}

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

func (kv *KeyValueStore) CreateUser(user User) int64 {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	user.ID = rand.Int63n(1000)
	kv.users[user.ID] = user
	return user.ID
}

func (kv *KeyValueStore) CreateOrder(order Order) int64 {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	order.ID = rand.Int63n(1000)
	kv.orders[order.ID] = order
	return order.ID
}

func (kv *KeyValueStore) LinkUserToOrder(userID, orderID int64) bool {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	_, ok := kv.users[userID]
	if !ok {
		return false
	}
	_, ok = kv.orders[orderID]
	if !ok {
		return false
	}
	if value, ok := kv.userOrders[userID]; ok {
		kv.userOrders[userID] = append(value, orderID)
	} else {
		kv.userOrders[userID] = []int64{orderID}
	}
	return true
}

func (kv *KeyValueStore) GetUserOrders(userID int64) []Order {
	kv.mu.RLock()
	defer kv.mu.RUnlock()
	orders := make([]Order, 0)
	if orderIDs, ok := kv.userOrders[userID]; ok {
		for _, orderID := range orderIDs {
			if order, ok := kv.orders[orderID]; ok {
				orders = append(orders, order)
			}
		}
	}
	return orders
}

var kvStore *KeyValueStore

func main() {
	kvStore = NewKeyValueStore("users.json", "orders.json", "userOrders.json")
	defer kvStore.save()

	http.HandleFunc("/create-user", createUserHandler)
	http.HandleFunc("/create-order", createOrderHandler)
	http.HandleFunc("/link-user-order", linkUserOrderHandler)
	http.HandleFunc("/get-user-orders", getUserOrdersHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid user data", http.StatusBadRequest)
		return
	}

	userID := kvStore.CreateUser(user)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int64{"userID": userID})
}

func createOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var order Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid order data", http.StatusBadRequest)
		return
	}

	orderID := kvStore.CreateOrder(order)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int64{"orderID": orderID})
}

func linkUserOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		UserID  int64 `json:"userID"`
		OrderID int64 `json:"orderID"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	success := kvStore.LinkUserToOrder(request.UserID, request.OrderID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"success": success})
}

func getUserOrdersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userIDStr := r.URL.Query().Get("userID")
	if userIDStr == "" {
		http.Error(w, "User ID parameter required", http.StatusBadRequest)
		return
	}

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid User ID", http.StatusBadRequest)
		return
	}

	orders := kvStore.GetUserOrders(userID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}
