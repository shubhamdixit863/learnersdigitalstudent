package main

import (
	"crud/internal/routes"
	"fmt"
	"log"
)

func main() {
	router := routes.SetUpRoutes()

	port := "8090"
	fmt.Println("Starting your HTTP server on port", port)
	log.Fatal(router.Run(":" + port))
}
