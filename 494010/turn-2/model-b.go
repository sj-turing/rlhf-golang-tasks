package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	// Initialize database
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/mydb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // PITFALL: Potential resource leak if db.Close() fails

	// Initialize RabbitMQ connection
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close() // PITFALL: Potential resource leak if conn.Close() fails

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close() // PITFALL: Potential resource leak if ch.Close() fails

	q, err := ch.QueueDeclare(
		"new_users", // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		log.Fatal(err)
	}

	// Create a new user and publish it to RabbitMQ
	user := User{ID: 1, Name: "John Doe"}
	if err := createUser(db, user); err != nil {
		log.Fatal(err)
	}
	if err := publishUser(ch, q, user); err != nil {
		log.Fatal(err)
	}

	fmt.Println("User created and published successfully!")
}

type User struct {
	ID   int
	Name string
}

func createUser(db *sql.DB, user User) error {
	stmt, err := db.Prepare("INSERT INTO users (id, name) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close() // PITFALL: Potential resource leak if stmt.Close() fails

	_, err = stmt.Exec(user.ID, user.Name)
	return err
}

func publishUser(ch *amqp.Channel, q amqp.Queue, user User) error {
	body, err := json.Marshal(user)
	if err != nil {
		return err
	}

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	return err
}
