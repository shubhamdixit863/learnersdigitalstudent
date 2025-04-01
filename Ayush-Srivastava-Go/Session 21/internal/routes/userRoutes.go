package routes

import (
	"userdb/internal/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes registers all application routes
func SetupRoutes(r *gin.Engine) {
	// v1 := r.Group("/api/v1")
	
	r.POST("/register", controllers.RegisterUser)
	r.GET("/get", controllers.GetUsers)
	r.PUT("/update", controllers.UpdateUser)
	r.DELETE("/delete", controllers.DeleteUser)
	
}
