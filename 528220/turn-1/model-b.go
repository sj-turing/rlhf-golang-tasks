package main

import (
	"fmt"
	"time"
)

// Error type for payment processing errors
type PaymentError struct {
	Code    string
	Message string
}

func (e PaymentError) Error() string {
	return fmt.Sprintf("Payment Error: %s - %s", e.Code, e.Message)
}

// Callback function type to handle payment notifications
type PaymentNotificationCallback func(paymentID string, amount float64, err error)

// Simulate an asynchronous payment processing function
func processPaymentAsync(paymentID string, amount float64, callback PaymentNotificationCallback) {
	go func() {
		time.Sleep(2 * time.Second) // Simulate asynchronous operation

		// Simulate a possible error condition
		if amount < 0 {
			callback(paymentID, amount, PaymentError{"InvalidAmount", "Amount must be positive."})
			return
		}

		callback(paymentID, amount, nil)
	}()
}

func main() {
	// Handle payment notifications asynchronously
	processPaymentAsync("12345", -100.00, func(paymentID string, amount float64, err error) {
		if err != nil {
			// Handle payment error
			fmt.Println("Payment Failed:", err)
		} else {
			// Handle successful payment
			fmt.Printf("Payment Successful: ID=%s, Amount=%.2f\n", paymentID, amount)
		}
	})

	// Continue doing other work...
	fmt.Println("Processing other tasks...")

	// Wait for a bit to allow the asynchronous payment processing to complete
	time.Sleep(3 * time.Second)
}
