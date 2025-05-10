package controllers

import (
	"log"
	"net/http"
	"strconv"
	"userdb/internal/models"
	"userdb/internal/utils"

	"github.com/gin-gonic/gin"
)

var users []models.User

// RegisterUser handles user registration
func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.InvalidRequestBodyError})
		return
	}


	user.ID = len(users) + 1
	users = append(users, user)

	log.Println(utils.UserRegisteredSuccess, user.Name)
	c.JSON(http.StatusCreated, user)
}

// GetUsers retrieves all users or a specific user by ID
func GetUsers(c *gin.Context) {
	idParam := c.Query(utils.SearchByIDParam)
	if idParam == "" { 
		c.JSON(http.StatusOK, users)
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.InvalidIDParamError})
		return
	}

	for _, user := range users {
		if user.ID == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": utils.UserNotFoundError})
}

// UpdateUser updates user details
func UpdateUser(c *gin.Context) {
	idParam := c.Query(utils.SearchByIDParam)
	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.InvalidIDParamError})
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.InvalidIDParamError})
		return
	}

	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.InvalidRequestBodyError})
		return
	}

	for index, user := range users {
		if user.ID == id {
			users[index].Name = updatedUser.Name
			users[index].Email = updatedUser.Email
			log.Println(utils.UserUpdatedSuccess, users[index].Name)
			c.JSON(http.StatusOK, users[index])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": utils.UserNotFoundError})
}

// DeleteUser removes a user by ID
func DeleteUser(c *gin.Context) {
	idParam := c.Query(utils.SearchByIDParam)
	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.InvalidIDParamError})
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.InvalidIDParamError})
		return
	}

	for i, user := range users {
		if user.ID == id {
			deletedUser := users[i]
			users = append(users[:i], users[i+1:]...)
			log.Println(utils.UserDeletedSuccess, user.Name)
			c.JSON(http.StatusOK, deletedUser)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": utils.UserNotFoundError})
}
