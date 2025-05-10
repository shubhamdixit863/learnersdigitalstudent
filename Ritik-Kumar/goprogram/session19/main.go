package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

// User structure
type User struct {
	Name   string `json:"name"`
	Id     int64  `json:"id"`
	Salary int64  `json:"salary"`
}

var users = make(map[int64]User) // Store users

func main() {
	// Home route
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.Method == "GET" {
			w.Write([]byte("hello world"))
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
	})

	// Register user
	http.HandleFunc("/register", func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		defer req.Body.Close()
		var user User
		body, err := io.ReadAll(req.Body)
		if err != nil || json.Unmarshal(body, &user) != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		users[user.Id] = user
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, "User registered successfully")
	})

	// Get users (all or by ID)
	http.HandleFunc("/get", func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		idStr := req.URL.Query().Get("id")
		if idStr == "" {
			json.NewEncoder(w).Encode(users) // Return all users
			return
		}

		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		user, exists := users[id]
		if !exists {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(user) // Return specific user
	})

	// Update user
	http.HandleFunc("/update", func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPut {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		idStr := req.URL.Query().Get("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		defer req.Body.Close()
		var user User
		body, err := io.ReadAll(req.Body)
		if err != nil || json.Unmarshal(body, &user) != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		users[id] = user
		fmt.Fprint(w, "User updated successfully")
	})

	// Delete user
	http.HandleFunc("/delete", func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodDelete {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		idStr := req.URL.Query().Get("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		if _, exists := users[id]; !exists {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		delete(users, id)
		fmt.Fprint(w, "User deleted successfully")
	})

	// Start the server
	log.Println("Server running on port 8080")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
