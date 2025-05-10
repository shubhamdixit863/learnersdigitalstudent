package server

import "errors"

// Slice to store user data
var users []User
var nextID = 1 // Auto-increment ID

// Add a new user
func AddUser(name, email string) User {
	user := User{ID: nextID, Name: name, Email: email}
	users = append(users, user)
	nextID++
	return user
}

// Get all users
func GetAllUsers() []User {
	return users
}

// Get user by ID
func GetUserByID(id int) (*User, error) {
	for _, user := range users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("User not found")
}

// Update user by ID
func UpdateUser(id int, name, email string) error {
	for i, user := range users {
		if user.ID == id {
			users[i].Name = name
			users[i].Email = email
			return nil
		}
	}
	return errors.New("User not found")
}

// Delete user by ID
func DeleteUser(id int) error {
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...) // Remove user
			return nil
		}
	}
	return errors.New("User not found")
}
