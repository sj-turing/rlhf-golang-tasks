package main

import (
	"fmt"
	"sync"
	"time"
)

type PurchaseOrder struct {
	OrderID  string
	Item     string
	Quantity int
}

var (
	orders         []PurchaseOrder
	ordersMutex    sync.Mutex // Mutex to protect access to orders slice
	maxOrders      = 10       // Maximum number of orders to store
	orderGenerator chan PurchaseOrder
	orderProcessor chan PurchaseOrder
	done           = make(chan bool)
)

func init() {
	orderGenerator = make(chan PurchaseOrder, 10)
	orderProcessor = make(chan PurchaseOrder, 10)
}

// GenerateOrder creates a new purchase order and sends it to the generator channel
func GenerateOrder() {
	for i := 0; i < 100; i++ {
		order := PurchaseOrder{
			OrderID:  fmt.Sprintf("O%06d", i+1),
			Item:     fmt.Sprintf("Item %d", i+1),
			Quantity: int(rand.Intn(100) + 1),
		}
		orderGenerator <- order
		time.Sleep(time.Millisecond * 10)
	}
	close(orderGenerator)
}

// ProcessOrder handles incoming orders from the processor channel
func ProcessOrder() {
	for order := range orderProcessor {
		ordersMutex.Lock()
		defer ordersMutex.Unlock()
		if len(orders) < maxOrders {
			orders = append(orders, order)
		} else {
			orders = orders[1:] // Remove the oldest order
			orders = append(orders, order)
		}
		fmt.Println("Processed Order:", order)
	}
	close(done)
}

func main() {
	go GenerateOrder()
	go ProcessOrder()

	<-done
	fmt.Println("\nLatest orders:")
	for _, order := range orders {
		fmt.Println(order)
	}
}
