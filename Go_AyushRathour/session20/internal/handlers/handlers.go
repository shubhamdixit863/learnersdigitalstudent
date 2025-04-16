package handlers

import (
	"crud/internal/model"
	"crud/internal/service"
	"crud/internal/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Initialize storage and service
var userStore = storage.NewUserStore()
var userService = service.NewUserService(userStore)

// RegisterUser handles user registration
func RegisterUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	registeredUser, _ := userService.RegisterUser(user)
	c.JSON(http.StatusCreated, registeredUser)
}

// GetUser function retrieves a user by ID or all users
func GetUser(c *gin.Context) {
	idParam := c.Query("id")

	if idParam == "" {
		c.JSON(http.StatusOK, userService.GetAllUsers())
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	user, err := userService.GetUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUser modifies an existing user
func UpdateUser(c *gin.Context) {
	idParam := c.Query("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedUser model.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	user, err := userService.UpdateUser(id, updatedUser)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUser removes a user by ID
func DeleteUser(c *gin.Context) {
	idParam := c.Query("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = userService.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
