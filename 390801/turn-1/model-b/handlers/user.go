package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"390801/turn-1/model-b/models"
	"390801/turn-1/model-b/services"

	"github.com/gorilla/mux"
)

func NewUserHandlers(userService *services.UserService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handleCreateUser(w, r, userService)
		case http.MethodGet:
			handleGetUser(w, r, userService)
		case http.MethodPut:
			//			handleUpdateUser(w, r, userService)
		case http.MethodDelete:
			//			handleDeleteUser(w, r, userService)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
}

func handleCreateUser(w http.ResponseWriter, r *http.Request, userService *services.UserService) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := userService.CreateUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func handleGetUser(w http.ResponseWriter, r *http.Request, userService *services.UserService) {
	allVars := mux.Vars(r)
	fmt.Printf("\n\nallVar = %v\n\n", allVars)

	requestId := allVars["id"]
	id, err := strconv.Atoi(requestId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := userService.GetUserByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// ... similar handlers for UpdateUser and DeleteUser
