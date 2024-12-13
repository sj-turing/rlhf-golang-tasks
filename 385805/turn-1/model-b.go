import (
	"database/sql"
	"fmt"
	"time"
)

func profileDatabaseTransaction(db *sql.DB, query string) error {
	start := time.Now()

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err := tx.Rollback(); err != nil {
			fmt.Println("Rollback failed:", err)
		}
		elapsed := time.Since(start)
		fmt.Printf("Query execution time: %s\n", elapsed)
	}()

	// Your database query here
	_, err = tx.Exec(query)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
