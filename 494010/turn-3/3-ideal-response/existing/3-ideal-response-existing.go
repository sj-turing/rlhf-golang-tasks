package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	PostgresHost   = "localhost"
	PostgresPort   = 5432
	PostgresUser   = "temp"
	PostgresPass   = "temp"
	PostgresDBName = "mydb"
)

func createAndPublishUser(ctx context.Context, db *sql.DB, username string) {
	conn, ch, err := createRabbitMQConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer closeRabbitMQ(conn, ch)

	query := `INSERT INTO users (username) VALUES ($1)`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, username)
	if err != nil {
		log.Println("Error creating user:", err)
		return
	}

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
	conn, err := amqp.Dial("amqp://username:password@localhost:5672/")
	if err != nil {
		fmt.Printf("\n\nerror while connecting to RMQ: %v", err)
		return nil, nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
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

func getPostgresConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		PostgresHost, PostgresPort, PostgresUser, PostgresPass, PostgresDBName)
}

func main() {
	ctx := context.Background()
	db, err := sql.Open("postgres", getPostgresConnectionString())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createAndPublishUser(ctx, db, "user-1")
}
