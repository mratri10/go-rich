package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mratri10/go-rich/models"
	"github.com/mratri10/go-rich/utils"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	hash, err := utils.HashPassword(input.Password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	if err := models.CreatedUser(input.Username, hash, input.Name, 0); err != nil {
		http.Error(w, "Username already exists", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	user, err := models.GetUserByUsernameForLogin(input.Username)

	if err != nil || !utils.CheckPasswordHash(input.Password, user.Password) {
		http.Error(w, "Invalid Username or Password", http.StatusUnauthorized)
		return
	}
	println(user.ID)
	token, err := utils.GenerateToken(user.ID, input.Username)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
