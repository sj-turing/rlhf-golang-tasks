package rabbitmq

import (
	"log"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Rmqer interface {
	Health() (string, error)
	Close() error
}

type Client struct {
	*amqp.Connection
	ch   *amqp.Channel
	mu   *sync.RWMutex
	conf *Config
}

func NewClient(c *Config) (Rmqer, error) {
	log.Println("Connecting to RMQ")
	conn, err := amqp.Dial(c.GetAMQPURL())
	if err != nil {
		log.Printf("New rmq conn failed: %v", err)
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	client := &Client{Connection: conn, ch: ch, mu: &sync.RWMutex{}, conf: c}

	return client, nil
}

// Health for rmq health check
func (c *Client) Health() (string, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if c.Connection.IsClosed() {
		return "rabbitMQ conn is down", amqp.ErrClosed
	}
	return "UP", nil
}

func (c *Client) Close() error {
	if c == nil {
		return nil
	}

	if err := c.ch.Close(); err != nil {
		return err
	}

	log.Printf("Closing the RMQ connection")
	return c.Connection.Close()
}
