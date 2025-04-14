package main

import (
	"assesment_5/internal/handlers"
	"assesment_5/internal/middlewares"
	"assesment_5/internal/repository"
	"assesment_5/internal/services"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const (
	env              = ".env"
	localhost        = "localhost:8080"
	envNotFoundError = "Cannot find env file"
)

func main() {
	// We will load the env file
	err := godotenv.Load(env)
	if err != nil {
		log.Println(envNotFoundError)
		return
	}
	r := gin.Default()

	//handler object
	repo := repository.NewMemoryRepo()
	jwtService := &services.JWTService{}
	handler := handlers.NewHandler(repo, jwtService)

	api := r.Group("/posts")
	api.GET("", handler.GetPosts)
	api.POST("", handler.CreatePost)
	api.PUT("/:id", middlewares.AuthorizationMiddleware(), handler.UpdatePost)
	api.DELETE("/:id", middlewares.AuthorizationMiddleware(), handler.DeletePost)

	err = r.Run(localhost)
	if err != nil {
		return
	}
}
