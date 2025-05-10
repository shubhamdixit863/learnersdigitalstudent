package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}
type Handler struct {
	users []User //dependency injection
}

func NewHandler(users []User) *Handler {
	return &Handler{users}
}

func (h *Handler) Create(c *gin.Context) {
	var user User

	err := c.BindJSON(&user)
	if err != nil {
		c.Error(errors.New("Error in reading the data"))
		return
	}
	id, err := uuid.NewUUID()
	if err != nil {
		c.Error(errors.New("Error in creating the user"))
		return
	}
	user.ID = id.String()

	//now we can insert it in the

	h.users = append(h.users, user)
	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
		"id":      user.ID,
	})
}

func (h *Handler) Get(c *gin.Context) {
	var users []map[string]string

	for _, user := range h.users {
		users = append(users, map[string]string{
			"id":   user.ID,
			"name": user.Name,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"users":   users,
	})
}

func (h *Handler) GetById(c *gin.Context) {
	id := c.Param("id")
	for _, user := range h.users {
		if user.ID == id {
			c.JSON(http.StatusOK, gin.H{
				"message": "Success",
				"user": map[string]string{
					"id":   user.ID,
					"name": user.Name,
				},
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "No user Found",
		"user":    nil,
	})
}

func (h *Handler) Update(c *gin.Context) {

	id := c.Param("id")
	var user User
	err := c.BindJSON(&user)
	if err != nil {
		c.Error(errors.New("Error in reading the data"))
		return
	}
	user.ID = id

	for i := 0; i < len(h.users); i++ {
		if h.users[i].ID == id {
			h.users[i] = user
			c.JSON(http.StatusOK, gin.H{
				"message": "User updated successfully",
				"user":    user,
			})
			return
		}
	}
}

func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")

	for i := 0; i < len(h.users); i++ {
		if h.users[i].ID == id {
			h.users = append(h.users[:i], h.users[i+1:]...) // Remove user
			c.JSON(http.StatusOK, gin.H{
				"message": "User deleted successfully",
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"message": "User not found",
	})
}
