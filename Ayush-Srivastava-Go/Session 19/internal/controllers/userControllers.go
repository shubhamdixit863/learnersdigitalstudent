package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"userdb/internal/models"
	"userdb/internal/utils"
)


var users []models.User

// Register a new user
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost { 
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(utils.InvalidRequestBodyError))
		return
	}

	var user models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(utils.InvalidRequestBodyError))
		return
	}

	user.ID = len(users) + 1
	users = append(users, user)

	w.Header().Add("Content-Type", "application/json")

	log.Println(utils.UserRegisteredSuccess, user.Name)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user) 
}


func GetUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	idParam := r.URL.Query().Get(utils.SearchByIDParam)
	if idParam == "" { 
		json.NewEncoder(w).Encode(users)
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(utils.InvalidIDParamError))
		return
	}

	w.Header().Add("Content-Type", "application/json")


	for _, user := range users {
		if user.ID == id {
			json.NewEncoder(w).Encode(user)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(utils.UserNotFoundError))
}

// Update a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut { 
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	idParam := r.URL.Query().Get(utils.SearchByIDParam)
	if idParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(utils.InvalidIDParamError))
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(utils.InvalidIDParamError))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(utils.InvalidRequestBodyError))
		return
	}

	var updatedUser models.User
	err = json.Unmarshal(body, &updatedUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(utils.InvalidRequestBodyError))
		return
	}

	
	w.Header().Add("Content-Type", "application/json")


	for index, user := range users {
		if user.ID == id {
			// users[index] = updatedUser
			users[index].Name = updatedUser.Name
			users[index].Email = updatedUser.Email
			log.Println(utils.UserUpdatedSuccess, users[index].Name)
			json.NewEncoder(w).Encode(users[index]) 
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(utils.UserNotFoundError))
}

// Delete a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	idParam := r.URL.Query().Get(utils.SearchByIDParam)
	if idParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(utils.InvalidIDParamError))
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(utils.InvalidIDParamError))
		return
	}

	w.Header().Add("Content-Type", "application/json")


	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			log.Println(utils.UserDeletedSuccess, user.Name)
			w.WriteHeader(http.StatusNoContent) 
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(utils.UserNotFoundError))
}
