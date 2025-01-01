package main

import (
	"fmt"
	"sync"
)

// Message represents a message sent through the broker.
type Message struct {
	Data  string
	Topic string
}

// MessageBroker defines the operations for interacting with a message broker.
type MessageBroker interface {
	Publish(message Message)
	Subscribe(topic string, callback func(message Message))
}

// InMemoryBroker implements the MessageBroker interface using an in-memory queue.
type InMemoryBroker struct {
	topics map[string][]func(message Message)
	mux    sync.Mutex
}

func NewInMemoryBroker() *InMemoryBroker {
	return &InMemoryBroker{
		topics: make(map[string][]func(message Message)),
	}
}

// Publish sends a message to the specified topic.
func (b *InMemoryBroker) Publish(message Message) {
	b.mux.Lock()
	defer b.mux.Unlock()

	// Find all subscribers for the message's topic
	if subscribers, ok := b.topics[message.Topic]; ok {
		for _, subscriber := range subscribers {
			subscriber(message)
		}
	}
}

// Subscribe registers a callback function for a specific topic.
func (b *InMemoryBroker) Subscribe(topic string, callback func(message Message)) {
	b.mux.Lock()
	defer b.mux.Unlock()

	// Append the callback to the list of subscribers for the topic
	b.topics[topic] = append(b.topics[topic], callback)
}

// ProcessMessage is a functional callback to process incoming messages.
func ProcessMessage(message Message) {
	fmt.Printf("Received message: %s\n", message.Data)
	// Simulate processing
	fmt.Println("Message processing completed.")
}

// AcknowledgeMessage is a functional callback to acknowledge message processing.
func AcknowledgeMessage(message Message) {
	fmt.Println("Acknowledged message processing.")
}

func main() {
	broker := NewInMemoryBroker()

	// Subscribe to a topic using functional callbacks
	broker.Subscribe("example-topic", func(message Message) {
		ProcessMessage(message)
		AcknowledgeMessage(message)
	})

	// Publish a message to the broker
	broker.Publish(Message{Data: "Hello, world!", Topic: "example-topic"})
}
