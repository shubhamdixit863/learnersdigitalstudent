package main

import (
	"log"
	"net/http"
	"practical/internal/handlers"

	"github.com/gin-gonic/gin"
)

func PingHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
func main() {
	r := gin.Default()
	r.LoadHTMLFiles("../internal/static/index.html")
	r.Static("/static", "../internal/static")
	r.GET("/", PingHandler)
	users := make([]handlers.User, 0)
	crudHandler := handlers.NewHandler(users)

	crudRoutes := r.Group("/api/v1")
	crudRoutes.POST("/create", crudHandler.Create)
	crudRoutes.GET("/get", crudHandler.Get)
	crudRoutes.GET("/get/:id", crudHandler.GetById)
	crudRoutes.PUT("/update/:id", crudHandler.Update)
	crudRoutes.DELETE("/delete/:id", crudHandler.Delete)
	log.Println("Server running on: http://localhost:8080")
	err := r.Run("localhost:8080")
	if err != nil {
		log.Println("Error starting server:", err)
	}
}
