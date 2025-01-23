package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/streadway/amqp"
)

// Notification represents a payment notification.
type Notification struct {
	ID     string  `json:"id"`
	Amount float64 `json:"amount"`
	Status string  `json:"status"`
}

// NotifyCallback defines the callback signature.
type NotifyCallback func(notification Notification, err error)

// Database configuration (using PostgreSQL as an example).
const (
	DbConnString = "user=youruser dbname=yourdb sslmode=disable password=yourpassword"
)

// RabbitMQ configuration.
const (
	RabbitMQURL  = "amqp://guest:guest@localhost:5672/"
	ExchangeName = "notifications"
	ExchangeType = "fanout"
)

// validateRequest validates the incoming notification request.
func validateRequest(notification Notification) error {
	if notification.ID == "" || notification.Amount <= 0 {
		return fmt.Errorf("invalid notification data")
	}
	return nil
}

// storeNotificationInDB stores the notification into the database.
func storeNotificationInDB(db *sql.DB, notification Notification) error {
	query := "INSERT INTO notifications (id, amount, status) VALUES ($1, $2, $3)"
	_, err := db.Exec(query, notification.ID, notification.Amount, notification.Status)
	return err
}

// publishNotification publishes the notification to RabbitMQ.
func publishNotification(notification Notification, callback NotifyCallback) {
	conn, err := amqp.Dial(RabbitMQURL)
	if err != nil {
		callback(notification, fmt.Errorf("failed to connect to RabbitMQ: %v", err))
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		callback(notification, fmt.Errorf("failed to open a channel: %v", err))
		return
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		ExchangeName, // name
		ExchangeType, // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		callback(notification, fmt.Errorf("failed to declare an exchange: %v", err))
		return
	}

	body, err := json.Marshal(notification)
	if err != nil {
		callback(notification, fmt.Errorf("failed to marshal notification: %v", err))
		return
	}

	err = ch.Publish(
		ExchangeName, // exchange
		"",           // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		callback(notification, fmt.Errorf("failed to publish message: %v", err))
		return
	}

	callback(notification, nil)
}

// notificationHandler handles incoming HTTP requests for notifications.
func notificationHandler(w http.ResponseWriter, r *http.Request) {
	var notification Notification
	if err := json.NewDecoder(r.Body).Decode(&notification); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := validateRequest(notification); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := sql.Open("postgres", DbConnString)
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	if err := storeNotificationInDB(db, notification); err != nil {
		http.Error(w, "Failed to store notification", http.StatusInternalServerError)
		return
	}

	callback := func(notification Notification, err error) {
		if err != nil {
			log.Printf("Error processing notification: %+v - Error: %v", notification, err)
			// Optionally, you can send another response to the client or take other actions.
			return
		}
		log.Printf("Successfully processed notification: %+v", notification)
	}

	go publishNotification(notification, callback)

	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintln(w, "Notification received and being processed")
}

func main() {
	http.HandleFunc("/notify", notificationHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
