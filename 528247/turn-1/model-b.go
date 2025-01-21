package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"math/rand"
	"sync"
	"time"
)

func fetchRecord(url string, wg *sync.WaitGroup) string {
	defer wg.Done()
	// Simulate fetching data from the URL
	time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
	return fmt.Sprintf("Result from URL: %s", url)
}
func insertRecord(record string, db *sql.DB, wg *sync.WaitGroup) {
	defer wg.Done()
	// Simulate inserting the record into the database
	stmt, err := db.Prepare("INSERT INTO records (record) VALUES ($1)")
	if err != nil {
		log.Fatalf("Error preparing statement: %v", err)
	}
	defer stmt.Close()

	if _, err = stmt.Exec(record); err != nil {
		log.Fatalf("Error inserting record: %v", err)
	}
}
func main() {
	db, err := sql.Open("postgres", "user=temp password=temp dbname=mydb sslmode=disable")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()
	// Create a slice of URLs
	urls := []string{"http://example.com/1", "http://example.com/2", "http://example.com/3"}
	wg := &sync.WaitGroup{}
	for _, url := range urls {
		wg.Add(2)
		// Launch goroutines to fetch and insert records
		go fetchRecord(url, wg)
		go insertRecord(fetchRecord(url, wg), db, wg)
	}
	wg.Wait()
	fmt.Println("All tasks completed.")
}
