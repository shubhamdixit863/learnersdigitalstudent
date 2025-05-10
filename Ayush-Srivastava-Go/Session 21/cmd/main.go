package main

import (
	"log"
	"os"
	"userdb/internal/routes"
	"userdb/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(utils.ENVFileLoadError)
		return
	}

	r := gin.Default()

	r.Static("/static", utils.StaticFolder)

	r.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	routes.SetupRoutes(r)

	address := os.Getenv("ENDPOINT_ADDRESS")
	if address == "" {
		address = ":8080" 
	}

	log.Println("Server running on PORT", address)
	if err := r.Run(address); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
