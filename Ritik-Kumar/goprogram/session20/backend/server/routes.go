package server

import "net/http"

func RegisterRoutes() {
	http.HandleFunc("/register", RegisterUser)
	http.HandleFunc("/get", GetUsers)
	http.HandleFunc("/getUser", GetUser)
	http.HandleFunc("/update", UpdateUserHandler)
	http.HandleFunc("/delete", DeleteUserHandler)
}
