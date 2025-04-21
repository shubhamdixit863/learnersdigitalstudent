package main

import (
	"JWT/internal/handlers"
	"JWT/internal/middlewares"
	"JWT/internal/repository"
	"JWT/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func main() {

	r := gin.Default()
	r.GET("/ping", PingHandler)
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//handler Object
	repo := repository.NewInMemory()
	jwtService := &services.JWTService{}
	handler := handlers.NewHandler(repo, jwtService)
	v1 := r.Group("/api/v1")
	auth := v1.Group("/auth")
	auth.POST("/signup", handler.Signup)
	auth.POST("/login", handler.Login)

	user := v1.Group("/user")
	user.GET("/getUsers", middlewares.AuthorizationMiddleware(), handler.GetAllUsers)

	err = r.Run("0.0.0.0:8090")
	if err != nil {
		return
	}

}
