package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/streadway/amqp"
)

// PaymentNotification represents a payment notification.
type PaymentNotification struct {
	ID      string  `json:"id"`
	Amount  float64 `json:"amount"`
	Status  string  `json:"status"`
	Message string  `json:"message"`
}

// PaymentRequest represents the request sent by partners.
type PaymentRequest struct {
	Amount float64 `json:"amount"`
	// Add other fields as needed
}

// PaymentResponse represents the response sent to partners.
type PaymentResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// NotifyCallback is a type for callback functions handling notifications.
type NotifyCallback func(notification PaymentNotification, err error)

// databaseClient provides an interface to a database.
type databaseClient interface {
	StorePaymentNotification(notification PaymentNotification) error
}

// rabbitMQClient provides an interface to RabbitMQ.
type rabbitMQClient interface {
	PublishEvent(notification PaymentNotification) error
}

// walletService handles payment notification API requests.
type walletService struct {
	dbClient databaseClient
	mqClient rabbitMQClient
}

// NewWalletService creates a new wallet service with the given dependencies.
func NewWalletService(dbClient databaseClient, mqClient rabbitMQClient) *walletService {
	return &walletService{
		dbClient: dbClient,
		mqClient: mqClient,
	}
}

// HandlePaymentNotification handles the payment notification API request.
func (ws *walletService) HandlePaymentNotification(w http.ResponseWriter, r *http.Request) {
	var req PaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate the request, e.g., check for valid amount, etc.
	if req.Amount <= 0 {
		http.Error(w, "Invalid payment amount", http.StatusBadRequest)
		return
	}

	notification := PaymentNotification{
		ID:     generateRandomID(), // Implement a function to generate a random ID
		Amount: req.Amount,
		Status: "Success", // You can set the status based on validation or other factors
	}

	go func() {
		if err := ws.handleNotificationAsync(notification); err != nil {
			log.Printf("Error handling notification: %v\n", err)
		}
	}()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(PaymentResponse{Success: true, Message: "Notification received"})
}

// handleNotificationAsync handles the payment notification asynchronously.
func (ws *walletService) handleNotificationAsync(notification PaymentNotification) error {
	if err := ws.dbClient.StorePaymentNotification(notification); err != nil {
		return fmt.Errorf("error storing notification: %w", err)
	}

	if err := ws.mqClient.PublishEvent(notification); err != nil {
		return fmt.Errorf("error publishing event: %w", err)
	}

	return nil
}

func main() {
	// Initialize database and RabbitMQ clients
	dbClient := NewMockDatabaseClient()
	mqClient := NewMockRabbitMQClient()

	ws := NewWalletService(dbClient, mqClient)

	http.HandleFunc("/notification", ws.HandlePaymentNotification)
	log.Println("Wallet Service listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
