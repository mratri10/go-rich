package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mratri10/go-rich/models"
	"github.com/mratri10/go-rich/utils"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("claims").(jwt.MapClaims)
	username, ok := claims["username"].(string)
	if !ok {
		http.Error(w, "Invalid Username in Token", http.StatusInternalServerError)
	}
	data, err := models.GetUserByUsername(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
	}
	json.NewEncoder(w).Encode(data)
}

func UserAllHandler(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("claims").(jwt.MapClaims)
	println(claims)
	data, err := models.GetUserAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(data)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("claims").(jwt.MapClaims)
	adminId, ok := claims["user_id"].(float64)
	if !ok {
		http.Error(w, "Invalid user_id in token", http.StatusUnauthorized)
		return
	}
	var input struct {
		Username string `json:"username"`
		Name     string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	hash, err := utils.HashPassword("Pass123")
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	if err := models.CreatedUser(input.Username, hash, input.Name, int(adminId)); err != nil {
		http.Error(w, "Username already exists", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User Added successfully"})
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("claims").(jwt.MapClaims)
	adminId, ok := claims["user_id"].(float64)
	if !ok {
		http.Error(w, "Invalid user_id in token", http.StatusUnauthorized)
	}
	var input struct {
		Password string `json:"password"`
		Username string `json:"username"`
		Name     string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	user, err := models.UpdateUserByUsername(input.Username, input.Name, input.Password, int(adminId))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}
	err = models.DeleteUseById(id)

	if err != nil {
		http.Error(w, "Delete User Failed", http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Data Deleted"})
}
