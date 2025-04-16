package services

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
	"os"
	"time"
)

type JWTService struct {
}

func (j *JWTService) GenerateToken(username string) (string, error) {

	//creating token with signing method and claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"nbf":      time.Now().Unix(),
	})

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", fmt.Errorf("JWT secret is not set")
	}
	return token.SignedString([]byte(secret))

}

//validate Token

func (j *JWTService) ValidateJWT(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return nil, fmt.Errorf("JWT secret is not set")
	}

	//parsing the token with a custom validation function
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Ensuring the signing method is exactly HS256
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		log.Println(claims["username"])
	} else {
		log.Println("token invalid")
	}

	if err != nil {
		log.Println("JWT validation failed", err)
		return nil, err
	}
	return token, nil
}
