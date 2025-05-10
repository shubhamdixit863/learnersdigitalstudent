package main

import (
	"1/middleware"
	"net/http"
)

func messageHandler(message string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(message))
	})
}

func main() {

	http.Handle("/", middleware.Middleware(messageHandler("Hello World")))
	http.ListenAndServe("localhost:8080", nil)
}
