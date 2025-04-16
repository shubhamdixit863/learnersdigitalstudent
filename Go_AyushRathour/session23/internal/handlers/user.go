package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetAllUsers(c *gin.Context) {
	users, err := h.repo.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed To Get All Users",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully Get All Users",
		"data":    users,
	})

}
