package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"session26/internal/services"
)

var (
	jwtServiceObj = services.JWTService{}
)

func AuthorizationMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		if authorization == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Please provide Authorization token",
			})
			c.Abort()
			return
		}

		_, err := jwtServiceObj.ValidateJWT(authorization)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Please provide Authorization token",
			})
			c.Abort()
			return
		}
		c.Next()
	}

}
