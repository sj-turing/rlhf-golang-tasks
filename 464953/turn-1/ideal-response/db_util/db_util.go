// db_util/db_util.go
package db_util

import (
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/go-sql-driver/mysql" // Replace this with your preferred database driver
)

const (
	dbDriver   = "mysql" // Replace with your preferred database driver
	dbUsername = "temp"
	dbPassword = "temp"
	dbName     = "mydb"
)

// DB represents the database connection
type DB struct {
	*sql.DB
}

// ConnectDB establishes a connection to the database
func ConnectDB() (*DB, error) {
	dsn := fmt.Sprintf("%s:%s@/%s", dbUsername, dbPassword, dbName)
	db, err := sql.Open(dbDriver, dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

// InsertParameters inserts URL query parameters into the database
func (db *DB) InsertParameters(params url.Values) error {
	stmt, err := db.Prepare("INSERT INTO parameters (some_key, some_value) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for key, values := range params {
		for _, value := range values {
			_, err := stmt.Exec(key, value)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// GetParameters retrieves URL query parameters from the database
func (db *DB) GetParameters() (url.Values, error) {
	params := make(url.Values)
	rows, err := db.Query("SELECT some_key, some_value FROM parameters")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var key, value string
		if err := rows.Scan(&key, &value); err != nil {
			return nil, err
		}
		params.Add(key, value)
	}

	return params, nil
}
