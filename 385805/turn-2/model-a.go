import (
	"database/sql"
	"fmt"
	"time"
)

func main() {
	// Open a database connection
	db, err := sql.Open("your_driver", "your_data_source_name")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	defer tx.Rollback() // Rollback in case of an error

	// Profile the transaction
	start := time.Now()

	// Your transaction code here
	_, err = tx.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", "John", "john@example.com")
	if err != nil {
		panic(err)
	}

	// Commit the transaction if no error occurs
	err = tx.Commit()
	if err != nil {
		panic(err)
	}

	// Calculate the execution time
	duration := time.Since(start)
	fmt.Printf("Transaction execution time: %s\n", duration)
}
