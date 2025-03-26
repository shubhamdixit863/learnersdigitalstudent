package routes

import (
	"net/http"
	"userdb/internal/controllers"
)

func SetupRoutes() {

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Welcome to UserDB"))
	// })
	http.HandleFunc("/register", controllers.RegisterUser)
	http.HandleFunc("/get", controllers.GetUsers)
	http.HandleFunc("/update", controllers.UpdateUser)
	http.HandleFunc("/delete", controllers.DeleteUser)
}