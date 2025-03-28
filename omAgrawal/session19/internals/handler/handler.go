package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"session19/internals/models"
	"session19/internals/services"
)

var newUserSl = services.NewUserSlice()

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Request body not supported"))
		}
		user := models.NewUser()
		err2 := json.Unmarshal(body, user)
		if err2 != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Request body not supported"))
		}
		newUserSl.RegisterUSer(user)
		d, _ := newUserSl.GetAllUser()
		fmt.Println(string(d))

	} else {
		w.WriteHeader(405)
	}
}

func HandleGetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		data, err := newUserSl.GetAllUser()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Request body not supported"))
			fmt.Println(err)
		}
		w.WriteHeader(200)
		w.Write(data)
	}
	w.WriteHeader(405)

}

func HandleGetById(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		u := r.URL.Query().Get("id")
		fmt.Println(u)

		data, err := newUserSl.GetUserById(u)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Request body not supported"))
			fmt.Println(err)
			return
		}
		w.WriteHeader(200)
		w.Write(data)
	}
	w.WriteHeader(405)

}

func HandleUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		id := r.URL.Query().Get("id")
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Request body not supported"))
		}
		user := models.NewUser()

		err2 := json.Unmarshal(body, user)
		fmt.Println(user)
		if err2 != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Request body not supported"))
		}
		newUserSl.UpdateUserById(id, user)
		w.WriteHeader(200)
		w.Write([]byte("successfully updated"))
	}
	w.WriteHeader(405)

}
func HandleDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		id := r.URL.Query().Get("id")
		err := newUserSl.DeleteUserById(id)
		fmt.Println(err)
		w.WriteHeader(200)
		w.Write([]byte("updateSuccesful"))
	}
	w.WriteHeader(405)

}
