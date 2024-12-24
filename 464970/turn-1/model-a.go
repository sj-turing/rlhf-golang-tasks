package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func queryDatabase(dbName string) error {
	// Connect to the database
	dsn := fmt.Sprintf("user:pass@tcp(127.0.0.1:3306)/%s?parseTime=true", dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("error opening database: %w", err)
	}
	defer db.Close() // Close the connection

	// Prepare the query
	query := "SELECT * FROM users"
	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparing query: %w", err)
	}
	defer stmt.Close() // Close the statement

	// Execute the query
	rows, err := stmt.Query()
	if err != nil {
		return fmt.Errorf("error querying database: %w", err)
	}
	defer rows.Close() // Close the rows result set

	// Process the results
	var userID int
	var username string
	for rows.Next() {
		err := rows.Scan(&userID, &username)
		if err != nil {
			return fmt.Errorf("error scanning row: %w", err)
		}
		fmt.Printf("User ID: %d, Username: %s\n", userID, username)
	}

	return nil // No errors
}

func main() {
	err := queryDatabase("my_database")
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
	}
}
