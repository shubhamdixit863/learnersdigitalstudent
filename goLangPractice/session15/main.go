package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	get, err := http.Get("http://127.0.0.1:8000/")
	if err != nil {
		return
	}

	body := get.Body
	all, err := io.ReadAll(body)
	if err != nil {
		return
	}

	log.Println("Data", string(all))
}
