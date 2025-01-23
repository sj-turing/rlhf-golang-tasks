package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/google/uuid"
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

type postgresDBMock struct{}

func NewMockDatabaseClient() *postgresDBMock {
	return &postgresDBMock{}
}

func (pdm *postgresDBMock) StorePaymentNotification(notification PaymentNotification) error {
	log.Println("storing record into database")
	time.Sleep(300 * time.Millisecond)
	if rand.Intn(100) <= 20 {
		return fmt.Errorf("Error while storing record into database")
	}
	return nil
}

// rabbitMQClient provides an interface to RabbitMQ.
type rabbitMQClient interface {
	PublishEvent(notification PaymentNotification) error
}

type eventPublisher struct{}

func NewMockRabbitMQClient() *eventPublisher {
	return &eventPublisher{}
}

func (ep *eventPublisher) PublishEvent(notification PaymentNotification) error {
	log.Println("publishing event into RMQ")
	time.Sleep(300 * time.Millisecond)
	if rand.Intn(100) <= 10 {
		return fmt.Errorf("Error while publishing event into RMQ")
	}
	return nil
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
		ID:     uuid.NewString(), // Implement a function to generate a random ID
		Amount: req.Amount,
		Status: "Success", // You can set the status based on validation or other factors
	}

	callback := func(notification PaymentNotification, err error) {
		if err != nil {
			log.Printf("Error processing notification: %+v - Error: %v", notification, err)
			// Optionally, you can send another response to the client or take other actions.
			return
		}
		log.Printf("Successfully processed notification: %+v", notification)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := ws.handleNotificationAsync(notification, callback); err != nil {
		log.Printf("Error handling notification: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(PaymentResponse{Success: false, Message: err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(PaymentResponse{Success: true, Message: "Notification received"})
}

// handleNotificationAsync handles the payment notification asynchronously.
func (ws *walletService) handleNotificationAsync(notification PaymentNotification, callback func(notification PaymentNotification, err error)) error {

	if err := ws.dbClient.StorePaymentNotification(notification); err != nil {
		callback(notification, fmt.Errorf("error storing notification: %w", err))
		return err
	}

	if err := ws.mqClient.PublishEvent(notification); err != nil {
		callback(notification, fmt.Errorf("error publishing event: %w", err))
		return err
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
