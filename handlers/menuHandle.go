package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mratri10/go-rich/models"
)

func AddMenu(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("claims").(jwt.MapClaims)
	adminId, ok := claims["user_id"].(float64)

	if !ok {
		http.Error(w, "Invalid when get User Id in Token", http.StatusInternalServerError)
		return
	}
	var input struct {
		Name   string `json:"name"`
		MenuId int    `json:"menuId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := models.CreateMenu(input.Name, int(adminId), input.MenuId); err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Menu Added Successfully"})
}
func GetMenu(w http.ResponseWriter, r *http.Request) {
	data, err := models.GetMenu()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(data)
}
func UpdateMenu(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("claims").(jwt.MapClaims)
	adminId, ok := claims["user_id"].(float64)
	if !ok {
		http.Error(w, "invalid when get User Id in Token", http.StatusInternalServerError)
		return
	}
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
	}
	var input struct {
		Name   string `json:"name"`
		MenuId int    `json:"menuId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	menu, err := models.UpdateMenu(id, input.Name, int(adminId), input.MenuId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(menu)
}

func DeleteMenu(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = models.DeleteMenu(id)
	if err != nil {
		http.Error(w, "Menu Failed when Delete", http.StatusBadGateway)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Menu Deleted Successfully"})
}

func AddRoleMenu(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("claims").(jwt.MapClaims)

	adminId, ok := claims["user_id"].(float64)
	if !ok {
		http.Error(w, "Invalid when get User Id in Token", http.StatusInternalServerError)
		return
	}
	var input struct {
		RoleId int `json:"roleId"`
		MenuId int `json:"menuId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	if err := models.CreateMenuRole(input.MenuId, input.RoleId, int(adminId)); err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User Setted Menu Successfully"})
}

func DeleteRoleMenu(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = models.DeleteMenuRole(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User Unsetted Menu Successfully"})
}
