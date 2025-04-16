package routes

import (
	"crud/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetUpRoutes() *gin.Engine {

	r := gin.Default()
	router := r.Group("/api/v1")
	router.POST("/register", handlers.RegisterUser)
	router.GET("/get", handlers.GetUser)
	router.PUT("/update", handlers.UpdateUser)
	router.DELETE("/delete", handlers.DeleteUser)

	router.Static("/home", "./internal/static")
	return r
}
