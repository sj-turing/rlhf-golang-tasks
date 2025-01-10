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
	wg              sync.WaitGroup
	purchaseOrders  = make([]PurchaseOrder, 0)
	orderQueue      = make(chan PurchaseOrder)
	processedOrders = make(chan PurchaseOrder)
)

func generatePurchaseOrder(id int) {
	defer wg.Done()
	for {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		customer := fmt.Sprintf("Customer %d", rand.Intn(1000))
		amount := rand.Float64() * 1000
		po := PurchaseOrder{OrderId: id, Customer: customer, Amount: amount}
		// Send the purchase order to the orderQueue channel for processing
		orderQueue <- po
	}
}
func processPurchaseOrders() {
	defer wg.Done()
	for {
		select {
		case po := <-orderQueue:
			// Process the purchase order
			fmt.Printf("Processing Purchase Order: %+v\n", po)
			processedOrders <- po
		default:
			// Simulate processing delay
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}
}
func main() {
	numGenerators := 5
	numProcessors := 3
	rand.Seed(time.Now().UnixNano())
	wg.Add(numGenerators + numProcessors)
	for i := 1; i <= numGenerators; i++ {
		go generatePurchaseOrder(i)
	}
	for i := 1; i <= numProcessors; i++ {
		go processPurchaseOrders()
	}
	go func() {
		defer wg.Done()
		for po := range processedOrders {
			// Acquire lock to safely access the shared slice
			purchaseOrders = append(purchaseOrders, po)
		}
	}()
	wg.Wait()
	close(processedOrders)
	// Print the final state of the purchase orders
	fmt.Println("Final Purchase Orders:")
	for _, po := range purchaseOrders {
		fmt.Printf("%+v\n", po)
	}
}
