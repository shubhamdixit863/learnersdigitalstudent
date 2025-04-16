package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"log/slog"
	"net/http"
	"session26/internal/dto"
	"session26/internal/models"
)

func (h *Handler) Signup(c *gin.Context) {
	var signupRequest dto.SignupRequest
	err := c.BindJSON(&signupRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Request",
			"error":   err.Error(),
		})
		return
	}

	//We check here if the user is already there
	_, err = h.repo.GetUserByUserName(c, signupRequest.Username)
	log.Println(err)
	if err == nil {
		c.JSON(http.StatusCreated, gin.H{
			"message": "User With Same Username already exists",
		})
		return
	}

	//converting dto to model
	userIdData, _ := uuid.NewUUID()
	userModel := models.User{
		Username:   signupRequest.Username,
		Password:   signupRequest.Password,
		FirstName:  signupRequest.FirstName,
		SecondName: signupRequest.SecondName,
		//Email:      signupRequest.Email,
		ID: userIdData.String(),
	}

	userId, err := h.repo.CreateUser(c, userModel)
	if err != nil {
		c.JSON(http.StatusCreated, gin.H{
			"message": "User Created",
			"ID":      userId,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "User Created",
		"ID":      userId,
	})
	return
}

func (h *Handler) Login(c *gin.Context) {
	var LoginRequest dto.LoginRequest
	err := c.BindJSON(&LoginRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Request",
		})
		return
	}
	log.Println("login request", LoginRequest)
	//checking if user is already there

	user, err := h.repo.GetUserByUserName(c, LoginRequest.Username)
	log.Println(err)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User Not Found",
		})

		return
	}

	// The jwt creation  code here

	jwt, err := h.jwtService.GenerateToken(user.Username)
	if err != nil {
		slog.Default().Error("Error in getting token", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed Generating the token",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Success",
		"Token":   jwt,
	})
	return
}
