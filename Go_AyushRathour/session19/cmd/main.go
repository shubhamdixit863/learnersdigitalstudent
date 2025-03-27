package main

import (
	"crud/internal/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := routes.SetupRoutes()

	port := "8080"
	fmt.Println("Starting HTTP server on port", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
