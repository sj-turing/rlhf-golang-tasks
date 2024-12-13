// main.go
package main

import (
	"390801/turn-2/model-a/repositories"
	"390801/turn-2/model-a/services"
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	var userRepo repositories.UserRepository
	dbType := os.Getenv("DB_TYPE") // "sql" or "nosql"

	if dbType == "nosql" {
		mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
		if err != nil {
			log.Fatal(err)
		}
		userRepo = repositories.NewNoSQLUserRepository(mongoClient, "appdb")
	} else {
		sqlDB, err := sql.Open("postgres", "user=username dbname=mydb sslmode=disable")
		if err != nil {
			log.Fatal(err)
		}
		userRepo = repositories.NewSQLUserRepository(sqlDB)
	}

	services.NewUserService(userRepo)

	// Further setup and HTTP server (handlers, routes)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
