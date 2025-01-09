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
	purchaseOrders := make([]PurchaseOrder, 0)

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

func generatePurchaseOrder(wg *sync.WaitGroup, purchaseOrders *[]PurchaseOrder) {
	defer wg.Done()
	for {
		customer := fmt.Sprintf("Customer %d", rand.Intn(1000))
		amount := rand.Float64() * 1000
		po := PurchaseOrder{OrderId: int(time.Now().UnixNano()), Customer: customer, Amount: amount}

		// Write access
		writeAccessChannel <- struct{}{}
		*purchaseOrders = append(*purchaseOrders, po)
		<-writeAccessChannel
	}
}

func processPurchaseOrders(wg *sync.WaitGroup, purchaseOrders *[]PurchaseOrder) {
	defer wg.Done()
	for {
		select {
		case <-time.After(time.Duration(rand.Intn(1000)) * time.Millisecond):
			// Read access
			if len(*purchaseOrders) > 0 {
				readAccessChannel <- struct{}{}
				po := (*purchaseOrders)[0]
				*purchaseOrders = (*purchaseOrders)[1:]
				<-readAccessChannel
				fmt.Printf("Processing Purchase Order: %+v\n", po)
			}
		}
	}
}

var writeAccessChannel = make(chan struct{}, 1)
var readAccessChannel = make(chan struct{}, numWorkers)
