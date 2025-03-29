package handlers

import (
	"crud/internal/model"
	"crud/internal/service"
	"crud/internal/storage"
	"encoding/json"
	"net/http"
	"strconv"
)

// Initialize storage and service
var userStore = storage.NewUserStore()
var userService = service.NewUserService(userStore)

// RegisterUser handles user registration
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	body := json.NewDecoder(r.Body)
	if err := body.Decode(&user); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	registeredUser, _ := userService.RegisterUser(user)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(registeredUser)
}

// GetUser function retrieves a user by ID or all users
func GetUser(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")

	if idParam == "" {
		json.NewEncoder(w).Encode(userService.GetAllUsers())
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	user, err := userService.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// UpdateUser modifies an existing user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedUser model.User
	body := json.NewDecoder(r.Body)
	if err := body.Decode(&updatedUser); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	user, err := userService.UpdateUser(id, updatedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// DeleteUser removes a user by ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = userService.DeleteUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
