package main

import (
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

// Define a Message struct as before
type Message struct {
    Data  string
    Topic string
}

// A Processor interface represents a function that takes a Message and returns a boolean indicating success
type Processor func(Message) bool

// Create a Consumer struct that holds the connection and channel to RabbitMQ, and a slice of processors
type Consumer struct {
	conn    *amqp.Connection
	ch      *amqp.Channel
	processors []Processor
}

// NewConsumer creates a new Consumer instance and initializes the RabbitMQ connection and channel
func NewConsumer(connectionString string, topics []string) (*Consumer, error) {
	conn, err := amqp.Dial(connectionString)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	// Declare the topics in RabbitMQ if they don't exist
	for _, topic := range topics {
		_, err := ch.QueueDeclare(topic, false, false, false, false, nil)
		if err != nil {
			ch.Close()
			conn.Close()
			return nil, err
		}
	}

	c := &Consumer{
		conn: conn,
		ch:   ch,
	}

	return c, nil
}

// AddProcessor adds a message processor to the consumer
func (c *Consumer) AddProcessor(p Processor) {
	c.processors = append(c.processors, p)
}

// ConsumeMessages starts the consumer and processes messages from the declared topics
func (c *Consumer) ConsumeMessages() {
	for _, topic := range c.topics {
		msgs, err := c.ch.Consume(
			topic, // queue
			"",    // consumer
			false, // auto-ack
			false, // exclusive
			false, // no-local
			false, // no-wait
			nil,   // args
		)
		if err != nil {
			log.Fatalf("Error consuming messages from topic %s: %v", topic, err)
		}

		go func() {
			for d := range msgs {
				message := Message{
					Data:  string(d.Body),
					Topic: d.RoutingKey,
				}

				// Process the message using each registered processor
				for _, processor := range c.processors {
					if success := processor(message); !success {
						// If any processor fails, nack the message and return
						d.Nack(false, true)
						return
					}
				}

				// If all processors succeed, ack the message
				d.Ack(false)
			}
		}()
	}
}

// Close closes the RabbitMQ connection and channel
func (c *Consumer) Close() {
	c.ch.Close()
	c.conn.Close()
}

// Example message processor
func processMessage(m Message) bool {
	fmt.Printf("Received message: %s from topic: %s\n", m.Data, m.Topic)
	
    // Simulate some processing time
    time.Sleep(time.Second * 2)
	
    // Return true if processing was successful, false otherwise
    return true
}

func main() {
	connectionString := "amqp://guest:guest@localhost:5672/"
	topics := []string{"example-topic"}
