package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type PurchaseOrder struct {
	OrderId  int
	Customer string
	Amount   float64
}

var (
	wg             sync.WaitGroup
	purchaseOrders []PurchaseOrder
	mutex          sync.Mutex
)

func generatePurchaseOrder(id int) {
	defer wg.Done()
	for {
		// Simulate random purchase order generation
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		customer := fmt.Sprintf("Customer %d", rand.Intn(1000))
		amount := rand.Float64() * 1000
		po := PurchaseOrder{OrderId: id, Customer: customer, Amount: amount}
		// Acquire lock to safely access the shared slice
		mutex.Lock()
		purchaseOrders = append(purchaseOrders, po)
		mutex.Unlock()
	}
}
func processPurchaseOrders() {
	defer wg.Done()
	for {
		// Simulate processing purchase orders after a random delay
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		// Acquire lock to safely access the shared slice
		mutex.Lock()
		if len(purchaseOrders) > 0 {
			po := purchaseOrders[0]
			purchaseOrders = purchaseOrders[1:]
			mutex.Unlock()
			fmt.Printf("Processing Purchase Order: %+v\n", po)
		} else {
			mutex.Unlock()
		}
	}
}
func main() {
	numGenerators := 5
	numProcessors := 3
	// Initialize random number generator
	rand.Seed(time.Now().UnixNano())
	// Start worker goroutines
	wg.Add(numGenerators + numProcessors)
	for i := 1; i <= numGenerators; i++ {
		go generatePurchaseOrder(i)
	}
	for i := 1; i <= numProcessors; i++ {
		go processPurchaseOrders()
	}
	// Wait for all goroutines to complete
	wg.Wait()
}
