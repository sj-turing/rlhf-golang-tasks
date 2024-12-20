package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("error while reading configuration: %v", err)
	}

	dbHost := viper.GetString("db.host")
	dbUser := viper.GetString("db.username")
	dbPass := viper.GetString("db.password")
	dbName := viper.GetString("db.name")
	dbConnString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPass, dbName)

	// Use a secure database driver to connect with the password
	db, err := sql.Open("postgres", dbConnString)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	fmt.Println("Successfully connected to the Database")
	defer db.Close()

	// Use prepared statements to execute queries with the password
	stmt, err := db.Prepare("SELECT * FROM users WHERE username = $1")
	if err != nil {
		log.Fatal("Error preparing the statement:", err)
	}

	dbUsername := "abc"

	defer stmt.Close()
	rows, err := stmt.Query(dbUsername)
	if err != nil {
		log.Fatal("Error executing the query:", err)
	}

	type User struct {
		ID       uint8
		Username string
	}

	for rows.Next() {
		user := User{}

		rows.Scan(&user.ID, &user.Username)

		fmt.Printf("User := %+v", user)
	}

	defer rows.Close()
}
