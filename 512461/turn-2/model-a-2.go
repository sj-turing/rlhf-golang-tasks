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

func main() {
	numWorkers := 10
	purchaseOrders := sync.Map{}

	var wg sync.WaitGroup

	// Generate purchase orders concurrently
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go generatePurchaseOrder(&wg, &purchaseOrders)
	}

	// Process purchase orders concurrently
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go processPurchaseOrders(&wg, &purchaseOrders)
	}

	wg.Wait()
}

func generatePurchaseOrder(wg *sync.WaitGroup, purchaseOrders *sync.Map) {
	defer wg.Done()
	for {
		customer := fmt.Sprintf("Customer %d", rand.Intn(1000))
		amount := rand.Float64() * 1000
		po := PurchaseOrder{OrderId: int(time.Now().UnixNano()), Customer: customer, Amount: amount}

		purchaseOrders.Store(po.OrderId, po)
	}
}

func processPurchaseOrders(wg *sync.WaitGroup, purchaseOrders *sync.Map) {
	defer wg.Done()
	for {
		select {
		case <-time.After(time.Duration(rand.Intn(1000)) * time.Millisecond):
			if k, v := purchaseOrders.LoadAndDelete(rand.Intn(100)); k != nil {
				fmt.Printf("Processing Purchase Order: %+v\n", v.(PurchaseOrder))
			}
		}
	}
}

