package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println()
		w.Write([]byte("Hello World"))
	})

	log.Println("Server is running on port: 8080")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Println("Error in starting the server", err)
		return
	}
}
