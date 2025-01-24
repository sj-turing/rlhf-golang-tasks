package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
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
	mu               *sync.RWMutex
	changeInFileChan chan struct{}
	users            map[int64]User
	orders           map[int64]Order
	userOrders       map[int64][]int64
	files            map[string]string
}

// NewKeyValueStore initializes the store and loads data
func NewKeyValueStore(userFile, orderFile, userOrderFile string) *KeyValueStore {
	kv := &KeyValueStore{
		users:            make(map[int64]User),
		orders:           make(map[int64]Order),
		userOrders:       make(map[int64][]int64),
		changeInFileChan: make(chan struct{}),
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

	go func() { kv.changeInFileChan <- struct{}{} }()
	w.Header().Set("Content-Type", "application/json")
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
	go func() { kv.changeInFileChan <- struct{}{} }()

	w.Header().Set("Content-Type", "application/json")
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
	go func() { kv.changeInFileChan <- struct{}{} }()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

// Retrieve all orders for a user

func (kv *KeyValueStore) GetUserOrders(w http.ResponseWriter, r *http.Request) {
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
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func (kv *KeyValueStore) changeInFileProcessor() {
	for range kv.changeInFileChan {
		kv.save()
	}
}

func main() {
	kvStore := NewKeyValueStore("users.json", "orders.json", "userOrders.json")
	go kvStore.changeInFileProcessor()

	svr := &http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/create-user", kvStore.CreateUser)
	http.HandleFunc("/create-order", kvStore.CreateOrder)
	http.HandleFunc("/link-user-order", kvStore.LinkUserOrder)
	http.HandleFunc("/get-user-orders", kvStore.GetUserOrders)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		fmt.Println("Server is starting at :8080")
		if err := svr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("error while closing the server: %v", err)
		}
	}()

	<-signalChan

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// graceful shutdown
	if err := svr.Shutdown(ctx); err != nil {
		log.Printf("Error while svr.Shutdown: %v", err)

		// in case of failure to graceful shutdown
		if errors.Is(err, context.DeadlineExceeded) {
			log.Printf("Executing force shutdown")
			// trying to forcefully shutdown
			if err := svr.Close(); err != nil {
				log.Printf("Server close failed: %v", err)
				log.Print("Proceeding for server exit")
			} else {
				log.Print("Server close completed")
			}
		}
	}

	close(signalChan)
	close(kvStore.changeInFileChan)
	kvStore.save()
	log.Print("The service terminated successfully with a graceful shutdown")
}
