package postgres

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type Connector interface {
	Get() *sql.DB
	Ping() error
	Close() error
}

type client struct {
	db *sql.DB
}

func NewConnection(config Config) (Connector, error) {

	db, err := sql.Open("postgres", config.formatDSN())
	if err != nil {
		log.Printf("Error while connecting to the postgres: ", err)
		return nil, err
	}

	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Second)

	return &client{
		db: db,
	}, nil
}

func (conn *client) Get() *sql.DB {
	return conn.db
}

func (conn *client) Close() error {
	log.Printf("Closing postgres connection")
	return conn.db.Close()
}

func (conn *client) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := conn.db.PingContext(ctx); err != nil {
		log.Printf("Error while ping: %v", err)
		return err
	}

	return nil
}
