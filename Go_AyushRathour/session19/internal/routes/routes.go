package routes

import (
	"crud/internal/handlers"
	"net/http"
)

func SetupRoutes() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/register", handlers.RegisterUser)
	router.HandleFunc("/get", handlers.GetUser)
	router.HandleFunc("/update", handlers.UpdateUser)
	router.HandleFunc("/delete", handlers.DeleteUser)

	fs := http.FileServer(http.Dir(".\\internal\\static"))
	router.Handle("/home/", http.StripPrefix("/home/", fs))

	return router
}
