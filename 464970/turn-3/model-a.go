package main

import (
	"fmt"
	"os"
)

func main() {
	// Example of managing a file handle
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // Ensure the file is closed after reading

	// Example of managing a database connection
	db, err := openDatabase()
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close() // Ensure the database connection is closed

	// Perform operations with the file and database
	fmt.Println("File opened successfully")
	fmt.Println("Database connection established")

	// Simulating an error to demonstrate defer execution
	err = someErrorousOperation()
	if err != nil {
		fmt.Println("Error occurred:", err)
	}
}

func openDatabase() (*databaseConnection, error) {
	fmt.Println("Opening database...")
	return &databaseConnection{}, nil
}

type databaseConnection struct{}

func (db *databaseConnection) Close() error {
	fmt.Println("Closing database connection...")
	return nil
}

func someErrorousOperation() error {
	// Simulate an error
	panic("This is a simulated error!")
}
