package main

import (
	"fmt"
	"sync"
)

type Message struct {
	Data string
}

type MessageBroker interface {
	Publish(message Message)
	Subscribe(topic string, callback func(message Message))
}

type InMemoryBroker struct {
	topics map[string][]func(message Message)
	mux    sync.Mutex
}

func NewInMemoryBroker() *InMemoryBroker {
	return &InMemoryBroker{
		topics: make(map[string][]func(message Message)),
	}
}

func (b *InMemoryBroker) Publish(message Message) {
	b.mux.Lock()
	defer b.mux.Unlock()

	// Find all subscribers for the message's topic
	if subscribers, ok := b.topics[message.Data]; ok {
		for _, subscriber := range subscribers {
			subscriber(message)
		}
	}
}

func (b *InMemoryBroker) Subscribe(topic string, callback func(message Message)) {
	b.mux.Lock()
	defer b.mux.Unlock()

	// Append the callback to the list of subscribers for the topic
	b.topics[topic] = append(b.topics[topic], callback)
}

type Application struct {
	broker MessageBroker
}

func NewApplication(broker MessageBroker) *Application {
	return &Application{
		broker: broker,
	}
}

func (app *Application) ProcessMessage(message Message) {
	fmt.Printf("Received message: %s\n", message.Data)
	// Do some processing here...

	// Acknowledge message processing by sending an acknowledgment back to the broker
	// (In this example, we don't have a broker, so we just print it)
	fmt.Println("Acknowledged message processing.")
}

func main() {
	broker := NewInMemoryBroker()
	app := NewApplication(broker)

	// Subscribe to a topic and register the application's message processing function as a callback
	broker.Subscribe("example-topic", app.ProcessMessage)

	// Publish a message to the broker
	broker.Publish(Message{Data: "Hello, world!"})
}
