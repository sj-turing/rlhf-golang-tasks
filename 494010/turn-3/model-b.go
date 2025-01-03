package main

import (
  "context"
  "database/sql"
  "fmt"
  "log"
  "sync"

  amqp "github.com/rabbitmq/amqp091-go"
  _ "github.com/lib/pq"
)

// common variables
const (
  dbConnString  = "user=dbuser dbname=dbpassword=host=localhost sslmode=disable"
  rabbitmqConnString = "amqp://guest:guest@localhost:5672/"
  publishQueue = "users.created"
)

// sync.WaitGroup to manage concurrent data handling
var wg sync.WaitGroup

func main() {
  ctx := context.Background()
  db, err := initDatabase()
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()

  conn, ch, err := createRabbitMQConnection()
  if err != nil {
    log.Fatal(err)
  }
  defer closeRabbitMQ(conn, ch)

  // Simulate user data
  users := []string{"user1", "user2", "user3", ...} // Large dataset

  wg.Add(len(users)) // Increase WaitGroup count for concurrent processing
  for _, username := range users {
    go processUser(ctx, db, ch, username)
  }
  wg.Wait() // Wait for all goroutines to complete
}

func initDatabase() (*sql.DB, error) {
  db, err := sql.Open("postgres", dbConnString)
  if err != nil {
    return nil, err
  }
  db.SetMaxOpenConns(100) // Set maximum open connections
  db.SetMaxIdleConns(20)  // Set maximum idle connections
  return db, nil
}

func createRabbitMQConnection() (*amqp.Connection, *amqp.Channel, error) {
  conn, err := amqp.Dial(rabbitmqConnString)
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
