package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"practical/internal/model"
	"practical/internal/services"
	"practical/internal/utils"
)

// func HomePage(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodPost {

// 	}
// }

func AddUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var user model.User
		err = json.Unmarshal(body, &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Println(user)
		err = services.AddUser(user)
		if err != nil {
			http.Error(w, utils.ServerErr, http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
		return
	}
	http.Error(w, utils.MethodErr, http.StatusMethodNotAllowed)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		users, err := services.GetUsers()
		if err != nil {
			http.Error(w, utils.ServerErr, http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
		return
	}
	http.Error(w, utils.MethodErr, http.StatusMethodNotAllowed)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, utils.QueryErr, http.StatusBadRequest)
			return
		}
		user, err := services.GetUserById(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
		return
	}
	http.Error(w, utils.MethodErr, http.StatusMethodNotAllowed)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, utils.QueryErr, http.StatusBadRequest)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, utils.BadReqErr, http.StatusBadRequest)
			return
		}
		var user model.User
		err = json.Unmarshal(body, &user)
		if err != nil {
			http.Error(w, utils.BadReqErr, http.StatusBadRequest)
			return
		}
		updatedUser, err := services.UpdateUser(id, user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updatedUser)
		return
	}
	http.Error(w, utils.MethodErr, http.StatusMethodNotAllowed)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, utils.QueryErr, http.StatusBadRequest)
			return
		}
		err := services.DeleteUser(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	}
	http.Error(w, utils.MethodErr, http.StatusMethodNotAllowed)
}
