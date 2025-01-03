package main

import (
	"log"

	"494010/turn-3/3-ideal-response/refactor/postgres"
	"494010/turn-3/3-ideal-response/refactor/rabbitmq"
)

func main() {
	db, err := postgres.NewConnection(postgres.Config{
		DBName:             "mydb",
		Hostname:           "localhost",
		Port:               5432,
		Username:           "temp",
		Password:           "temp",
		MaxOpenConnections: 5,
		MaxIdleConnections: 5,
	})
	if err != nil {
		log.Fatalf("Error while getting postgres connection: %v", err)
	}
	defer db.Close()

	rmq, err := rabbitmq.NewClient(&rabbitmq.Config{
		Host:     "localhost",
		Username: "username",
		Password: "password",
		Port:     5672,
	})
	if err != nil {
		log.Fatalf("Error while getting RMQ client: %v", err)
	}
	defer rmq.Close()

	status, err := rmq.Health()
	if err != nil {
		log.Fatalf("Error while checking the RMQ health: %v", err)
	}

	log.Printf("rmq status: %v", status)

	if err := db.Ping(); err != nil {
		log.Fatalf("Error while ping the database: %v", err)
	}

	log.Printf("Pigged to the database")
}
