package handlers

import (
	"net/http"

	"1/internal/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllUsers(c *gin.Context) {

	users, err := h.repo.GetAllUsers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed To Get users",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    users,
	})

}

func (h *Handler) DeleteUser(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Request",
		})
		return
	}
	err = h.repo.DeleteUser(c, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Couldnt delete user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted",
	})

}

func (h *Handler) UpdateUser(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Request",
		})
		return
	}
	err = h.repo.UpdateUser(c, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Couldnt update user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User Updated",
	})

}
