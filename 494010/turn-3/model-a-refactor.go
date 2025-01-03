package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"
	"sync/atomic"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
	"github.com/rabbitmq/amqp091-go"
)

var (
	dbPool     *sql.DB
	rabbitPool *redis.Client
)

func init() {
	ctx := context.Background()

	dbCfg := &sql.DBConfig{
		MaxOpenConns:    20,
		MaxIdleConns:    10,
		ConnMaxLifetime: 30 * time.Second,
	}
	db, err := sql.OpenWithConfig("postgres", dbCfg, "user=dbuser dbname=dbpassword=host=localhost sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	dbPool = db

	rabbitCfg := &redis.Options{
		Addrs:        []string{"localhost:6379"},
		PoolSize:     10,
		MinIdleConns: 1,
	}
	rabbitPool, err = redis.NewClient(rabbitCfg)
	if err != nil {
		log.Fatal(err)
	}
}

func createAndPublishUser(ctx context.Context, db *sql.DB, username string) {
	ch, err := rabbitPool.TxPipeline()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	query := `INSERT INTO users (username) VALUES (?)`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, username)
	if err != nil {
		log.Println("Error creating user:", err)
		return
	}

	err = ch.RPush("users.created", fmt.Sprintf(`{"username": "%s"}`, username)).Err()
	if err != nil {
		log.Println("Error publishing message:", err)
	}

	_, err = ch.Exec()
	if err != nil {
		log.Println("Error executing pipeline:", err)
	}
}

func main() {
	ctx := context.Background()
	var wg sync.WaitGroup
	const concurrency = 50

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(id int) {
			createAndPublishUser(ctx, dbPool, fmt.Sprintf("user%d", id))
			wg.Done()
		}(i)
	}

	go func() {
		for i := 0; i < concurrency; i++ {
			wg.Add(1)
			go func() {
				createAndPublishUser(ctx, dbPool, fmt.Sprintf("user%d", atomic.AddInt32(&i, 1000)))
				wg.Done()
			}()
		}
	}()

	wg.Wait()
}
