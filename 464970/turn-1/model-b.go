package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// Global variable
var db *sql.DB

func init() {
	// Initialize the database connection pool
	var err error
	db, err = sql.Open("mysql", "user:password@tcp(localhost:3306)/testdb")
	if err != nil {
		panic(err.Error())
	}
}
func main() {
	//Close the database connection pool after the main function has finished executing
	defer db.Close()
	result := getUserName(1)
	fmt.Println(result)
}

func getUserName(userId int) string {
	defer db.Close()
	// prepare the SQL query string
	stmt, err := db.Prepare("SELECT name FROM users WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	//Defer the closure of the statement
	defer stmt.Close()
	// Variable to store the result
	var username string
	// Execute the prepared statement and store the result in username
	err = stmt.QueryRow(userId).Scan(&username)
	if err != nil {
		panic(err.Error())
	}
	//Return the result
	return username
}
