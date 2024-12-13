package main

import (
	"log"
	"net/http"

	"390801/turn-1/model-b/handlers"
	"390801/turn-1/model-b/repositories"
	"390801/turn-1/model-b/services"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	userRepository := repositories.NewInMemoryUserRepository()
	userService := services.NewUserService(userRepository)

	r.Handle("/user", handlers.NewUserHandlers(userService))

	log.Fatal(http.ListenAndServe(":8080", r))
}
