package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"session22/internal/handlers"
	"session22/internal/middlewares"
)

func main() {
	mux := http.NewServeMux()
	fmt.Println(reflect.TypeOf(mux))
	mux.HandleFunc("/files", handlers.FileListHandler)
	mux.HandleFunc("/file", handlers.FileDataHandler)

	loggedMux := middleware.LogMiddleware(mux)
	fmt.Println("Server is running on port 8081...")
	log.Fatal(http.ListenAndServe(":8081", loggedMux))
}
