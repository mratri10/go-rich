package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mratri10/go-rich/models"
)

func AddRole(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("claims").(jwt.MapClaims)
	adminId, ok := claims["user_id"].(float64)
	if !ok {
		http.Error(w, "Invalid when get User Id in Token", http.StatusInternalServerError)
		return
	}
	var input struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	if err := models.CreatedRole(input.Name, int(adminId)); err != nil {
		http.Error(w, "Role already exists", http.StatusBadGateway)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Role Added Successfully"})
}
func GetRole(w http.ResponseWriter, r *http.Request) {
	data, err := models.GetRole()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(data)
}

func UpdateRole(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("claims").(jwt.MapClaims)
	adminId, ok := claims["user_id"].(float64)
	if !ok {
		http.Error(w, "Invalid when get User Id in Token", http.StatusInternalServerError)
		return
	}
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
	}
	var input struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}
	role, err := models.UpdateRole(id, input.Name, int(adminId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(role)
}
func DeleteRole(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = models.DeleteRole(id)

	if err != nil {
		http.Error(w, "Role Failed when Delete", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Role Deleted Successfully"})
}

func AddUserRole(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("claims").(jwt.MapClaims)

	adminId, ok := claims["user_id"].(float64)
	if !ok {
		http.Error(w, "Invalid when get User Id in Token", http.StatusInternalServerError)
		return
	}
	var input struct {
		UserId int `json:"userId"`
		RoleId int `json:"roleId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	if err := models.CreateRoleUser(input.RoleId, input.UserId, int(adminId)); err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User Setted Role Successfully"})
}
func DeleteUserRole(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = models.DeleteRoleUser(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User Unsetted Role Successfully"})
}
