package main

import (
	"fmt"
	"time"
)

// PaymentNotification represents a payment notification.
type PaymentNotification struct {
	ID      string
	Amount  float64
	Status  string
	Message string
}

// NotifyCallback is a type for callback functions handling notifications.
type NotifyCallback func(notification PaymentNotification, err error)

// asyncPaymentNotification simulates an asynchronous payment notification.
func asyncPaymentNotification(callback NotifyCallback) {
	go func() {
		// Simulate a delay in receiving the notification.
		time.Sleep(2 * time.Second)

		// Simulate a successful notification.
		notification := PaymentNotification{
			ID:     "abc123",
			Amount: 100.0,
			Status: "Success",
		}

		// Simulate an error condition.
		var err error
		if notification.Amount <= 0 {
			err = fmt.Errorf("invalid payment amount")
		}

		// Execute the callback with the notification and error (if any).
		callback(notification, err)
	}()
}

func main() {
	callback := func(notification PaymentNotification, err error) {
		if err != nil {
			fmt.Printf("Error handling notification: %v\n", err)
			// Log the error or take further action as needed.
			return
		}
		fmt.Printf("Received notification: %+v\n", notification)
		// Process the notification further as needed.
	}

	asyncPaymentNotification(callback)

	// Keep the main function alive to allow async operations to complete.
	time.Sleep(3 * time.Second)
}
