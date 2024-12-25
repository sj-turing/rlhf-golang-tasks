// main.go
package main

import (
	"fmt"
	"log"
	"net/url"

	"464953/turn-1/model-b/db_util"
)

func main() {
	db, err := db_util.ConnectDB()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	// Example URL with special characters in query parameters
	exampleURL := "https://example.com?name=Alice&city=New+York&hobby=coding&lang=Go%2B"
	u, err := url.Parse(exampleURL)
	if err != nil {
		log.Fatalf("Error parsing URL: %v", err)
	}

	// Insert parameters into the database
	if err := db.InsertParameters(u.Query()); err != nil {
		log.Fatalf("Error inserting parameters: %v", err)
	}

	// Retrieve parameters from the database
	retrievedParams, err := db.GetParameters()
	if err != nil {
		log.Fatalf("Error retrieving parameters: %v", err)
	}

	// Display retrieved parameters
	fmt.Println("Retrieved Parameters:")
	for key, values := range retrievedParams {
		for _, value := range values {
			fmt.Printf("%s: %s\n", key, value)
		}
	}
}
