package main

import (
	"fmt"
	"log"
	"net/http"
)

// Define a global slice to store user data
var users []map[string]string

// Handler function
func handler(w http.ResponseWriter, r *http.Request) {
	// Home Page
	if r.URL.Path == "/" {
		w.Write([]byte("Welcome to My Page!"))
		return
	}

	switch r.URL.Path {
	case "/register":
		if r.Method == "POST" {
			// Dummy user registration
			users = append(users, map[string]string{"id": "1", "name": "John Doe"})
			w.Write([]byte("User registered successfully"))
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	case "/get":
		w.Write([]byte(fmt.Sprintf("Users: %v", users)))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Page Not Found"))
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
