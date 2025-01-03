package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func createAndPublishUser(ctx context.Context, db *sql.DB, username string) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close() // Ensure RabbitMQ connection is closed

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close() // Ensure RabbitMQ channel is closed

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

func main() {
	ctx := context.Background()
	db, err := sql.Open("postgres", "user=dbuser dbname=dbpassword=host=localhost sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // Ensure the database connection is closed

	createAndPublishUser(ctx, db, "newuser")
}
