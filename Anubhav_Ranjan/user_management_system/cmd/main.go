package main

import (
	"fmt"
	"net/http"
	db "user_management_system/config"
	"user_management_system/internal/controller"
	"user_management_system/internal/middleware"
	"user_management_system/internal/repository"
	"user_management_system/internal/service"
)

func main() {
	// Initialize database connection
	db.InitializeDatabase()

	// Create repository, service, and controller
	userRepo := repository.NewUserRepository(db.GetDB())
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	// Routes
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			userController.GetAllUsers(w, r)
		} else if r.Method == http.MethodPost {
			userController.CreateUser(w, r)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			userController.GetUser(w, r)
		case http.MethodPut:
			userController.UpdateUser(w, r)
		case http.MethodDelete:
			userController.DeleteUser(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	loggedRoutes := middleware.LoggingMiddleware(http.DefaultServeMux)

	// Start the server
	fmt.Println("Server is running on port 8081")
	if err := http.ListenAndServe(":8081", loggedRoutes); err != nil {
		fmt.Println("Error Starting Server:", err)
	}
}
