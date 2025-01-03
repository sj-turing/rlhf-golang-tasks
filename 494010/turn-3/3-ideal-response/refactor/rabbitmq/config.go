package rabbitmq

import "fmt"

type Config struct {
	Username, Password, Host string
	Port                     int
}

// GetAMQPURL get rabbitMQ amqp url string
func (c *Config) GetAMQPURL() string {
	// "amqp://guest:guest@localhost:5672/"
	return fmt.Sprintf("amqp://%s:%s@%s:%d/", c.Username, c.Password, c.Host, c.Port)
}
