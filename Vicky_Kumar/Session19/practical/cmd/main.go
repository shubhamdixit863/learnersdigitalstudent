package main

import (
	"fmt"
	"net/http"
	"practical/internal/handlers"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("internal/static")))
	http.HandleFunc("/register", handlers.AddUser)
	http.HandleFunc("/get", handlers.GetUser)
	http.HandleFunc("/update", handlers.UpdateUser)
	http.HandleFunc("/delete", handlers.DeleteUser)
	http.HandleFunc("/getall", handlers.GetUsers)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
