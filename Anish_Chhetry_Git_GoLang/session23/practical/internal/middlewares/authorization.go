package middlewares

import (
	"log"
	"net/http"
	"practical/services"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	jwtServiceObj = services.JWTService{}
)

func AuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		log.Println("🔐 Authorization header:", authorization)

		if len(authorization) == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Please provide Authorization token",
			})
			c.Abort()
			return
		}

		// Strip "Bearer " prefix
		tokenString := strings.TrimPrefix(authorization, "Bearer ")

		_, err := jwtServiceObj.ValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Please provide valid Authorization token",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
