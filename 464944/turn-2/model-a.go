package main

import (
	"database/sql"
	"fmt"
	"io"
	"os"
)

// CleanupFile protocol for file cleanup
type CleanupFile interface {
	Close() error
}

// CleanupDatabase protocol for database connection cleanup
type CleanupDatabase interface {
	Close() error
}

// FileCloser implements the CleanupFile protocol
type FileCloser struct {
	*os.File
}

func (fc *FileCloser) Close() error {
	return fc.File.Close()
}

// DatabaseConnection implements the CleanupDatabase protocol
type DatabaseConnection struct {
	*sql.DB
}

func (db *DatabaseConnection) Close() error {
	return db.DB.Close()
}

func main() {
	// Open a file
	file, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	fc := &FileCloser{file}

	// Create a database connection
	db, err := sql.Open("driverName", "dataSourceName")
	if err != nil {
		panic(err)
	}
	defer db.Close() // Handle database connection cleanup with defer

	// Use defer for file cleanup
	defer func() {
		if err := fc.Close(); err != nil {
			fmt.Println("Error closing file:", err)
		}
	}()

	// Now you can use the file and the database connection as needed
	_, err = io.Copy(os.Stdout, fc)
	if err != nil {
		panic(err)
	}

	// Example operation with the database connection
	rows, err := db.Query("SELECT * FROM your_table")
	if err != nil {
		panic(err)
	}
	defer rows.Close() // Defer to close rows

	for rows.Next() {
		// Process each row
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}
}
