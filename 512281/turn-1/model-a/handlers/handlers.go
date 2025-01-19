package handlers

import (
	"512281/turn-1/model-a/repository"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	Repo repository.UserRepository
}

func (uh *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	user, err := uh.Repo.GetUserByUsername(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

// Similar HTTP handlers for CreateUser, UpdateUser, DeleteUser
