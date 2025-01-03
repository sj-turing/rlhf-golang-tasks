package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func createAndPublishUser(ctx context.Context, db *sql.DB, username string) {
	conn, ch, err := createRabbitMQConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer closeRabbitMQ(conn, ch) // Ensure RabbitMQ connection and channel are closed

	// SQL to create a user
	query := `INSERT INTO users (username) VALUES (?)`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // Ensure the prepared statement is closed

	_, err = stmt.ExecContext(ctx, username)
	if err != nil {
		log.Println("Error creating user:", err)
		return
	}

	// Publish message to RabbitMQ
	err = ch.Publish(
		"",              // exchange
		"users.created", // routing key
		false,           // mandatory
		false,           // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(fmt.Sprintf(`{"username": "%s"}`, username)),
		},
	)
	if err != nil {
		log.Println("Error publishing message:", err)
	}
}

func createRabbitMQConnection() (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close() // Close connection if channel creation fails
		return nil, nil, err
	}

	return conn, ch, nil
}

func closeRabbitMQ(conn *amqp.Connection, ch *amqp.Channel) {
	if ch != nil {
		err := ch.Close()
		if err != nil {
			log.Println("Error closing RabbitMQ channel:", err)
		}
	}
	if conn != nil {
		err := conn.Close()
		if err != nil {
			log.Println("Error closing RabbitMQ connection:", err)
		}
	}
}

func main() {
	ctx := context.Background()
	db, err := sql.Open("postgres", "user=dbuser dbname=dbpassword=host=localhost sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // Ensure the database connection is closed

	createAndPublishUser(ctx, db, "newuser")
}
