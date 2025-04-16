package main

import (
	"flag"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-mysql-org/go-mysql/driver"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
	"session26/internal/handlers"
	"session26/internal/middlewares"
	"session26/internal/repository"
	"session26/internal/services"
)

const dsn = "root:root@localhost:3306?myapp"

func mongoConnect(uri string) (*mongo.Client, error) {
	client, err := mongo.Connect(options.Client().
		ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	return client, nil
}

func main() {

	var dbtype string
	flag.StringVar(&dbtype, "dbtype", "mongodb", "Database type")
	flag.Parse()

	r := gin.Default()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}
	r.Use(cors.Default())

	var repo repository.DbRepository
	if dbtype == "mongodb" {
		client, err := mongoConnect("mongodb://localhost:27017/users")
		if err != nil {
			log.Fatal("Errorr connecting mongodb", err)
		}
		repo = repository.NewMongodb(client)

	} else if dbtype == "mysql" {

		db, err := sqlx.Connect("mysql", dsn)
		if err != nil {
			log.Println("Error connecting to mysql", err)
			return
		}
		repo = repository.NewMysqlReqo(db)

	} else if dbtype == "mysqlorm" {

		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			log.Println(err)
			return
		}
		repo = repository.NewMysqlOrm(db)

	} else if dbtype == "inmemory" {
		repo = repository.NewInMemory()
	}

	jwtService := &services.JWTService{}
	handler := handlers.NewHandler(repo, jwtService)
	v1 := r.Group("/api/v1")
	v1.GET("/healthz", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "All good",
		})
	})
	auth := v1.Group("/auth")
	auth.POST("/signup", handler.Signup)
	auth.POST("/login", handler.Login)

	//user routes
	user := v1.Group("/user")
	user.GET("/getUsers", middlewares.AuthorizationMiddleware(), handler.GetAllUsers)
	user.PUT("/updateUser", middlewares.AuthorizationMiddleware(), handler.UpdateUser)
	user.DELETE("/deleteUser", middlewares.AuthorizationMiddleware(), handler.DeleteUser)

	err = r.Run("localhost:8090")
	if err != nil {
		return
	}

}
