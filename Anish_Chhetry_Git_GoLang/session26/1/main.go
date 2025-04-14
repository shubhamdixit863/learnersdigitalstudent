package main

import (
	"1/internal/handlers"
	"1/internal/middlewares"
	"1/internal/repository"
	"1/internal/services"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-mysql-org/go-mysql/driver"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func main() {
	dsn := "root:root@localhost:3307?userdb"

	// We will load the env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Cannot find env file")
		return
	}

	r := gin.Default()

	r.Use(cors.Default())

	// Handler Object
	//repo := repository.NewInMemory()
	// sqlx connection
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Println("Error connecting to the db", err)
		return
	}
	repo := repository.NewMysqlReqo(db)
	jwtService := &services.JWTService{}
	handler := handlers.NewHandler(repo, jwtService)
	v1 := r.Group("/api/v1")
	v1.GET("/healthz", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "All good",
		})
	})
	// Two types of groups // auth routes
	auth := v1.Group("/auth") // /api/v1/auth

	auth.POST("/signup", handler.Signup)
	auth.POST("/login", handler.Login)

	// user routes
	user := v1.Group("/user") // /api/v1/user
	user.GET("/getUsers", middlewares.AuthorizationMiddleware(), handler.GetAllUsers)
	user.DELETE("/delete", middlewares.AuthorizationMiddleware(), handler.DeleteUser)
	user.PUT("/update", middlewares.AuthorizationMiddleware(), handler.UpdateUser)

	err = r.Run("localhost:8080")
	if err != nil {
		return
	}

}
