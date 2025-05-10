package handlers

import (
	"log"
	"log/slog"
	"net/http"
	"practical/internal/dto"
	"practical/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// https://stackoverflow.com/questions/3853749/what-is-the-difference-between-an-mvc-model-object-a-domain-object-and-a-dto
func (h *Handler) Signup(c *gin.Context) {

	var signupRequest dto.SignupRequest
	err := c.BindJSON(&signupRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Request",
		})
		return
	}

	// We check here if the user is already there
	_, err = h.repo.GetUserByUserName(signupRequest.Username)
	log.Println(err)
	if err == nil {
		c.JSON(http.StatusCreated, gin.H{
			"message": "User With Same Username already exists",
		})
		return
	}

	// What you have to do is save the data
	// We have to convert the dto to the model
	userIdData, _ := uuid.NewUUID()
	userModel := models.User{
		Username: signupRequest.Username,
		Password: signupRequest.Password,
		Name:     signupRequest.Name,
		Email:    signupRequest.Email,
		ID:       userIdData.String(),
	}

	userId := h.repo.CreateUser(userModel)
	c.JSON(http.StatusCreated, gin.H{
		"message": "User Created",
		"ID":      userId,
	})
	return

}

func (h *Handler) Login(c *gin.Context) {
	var loginRequest dto.LoginRequest
	err := c.BindJSON(&loginRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Request",
		})
		return
	}

	user, err := h.repo.GetUserByUserName(loginRequest.Username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User Not Found",
		})
		return
	}

	if user.Password != loginRequest.Password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid password",
		})
		return
	}

	jwt, err := h.jwtService.CreateJWT(user.Username)
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
}
