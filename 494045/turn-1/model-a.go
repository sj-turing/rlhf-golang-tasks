package main

import (
	"amqp"
	"log"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect: %s\n", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open channel: %s\n", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello",
		false, // durable
		false, // exclusive
		false, // auto-delete
		false, // no-wait
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare queue: %s\n", err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("Failed to register consumer: %s\n", err)
	}

	for d := range msgs {
		log.Printf("Received message: %s\n", string(d.Body))
	}
}
