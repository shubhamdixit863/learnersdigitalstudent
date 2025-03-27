package main

import (
	"log"
	"net/http"
	"practical/internal/services"
)

const (
	rootPath     = "/"
	registerPath = "/register"
	getPath      = "/get"
	updatePath   = "/update"
	deletePath   = "/delete"
	localHost    = "localhost:8080"
	serverLink   = "Server running on: http://localhost:8080"
	staticPath   = "../internal/static"
	stringStatic = "/static/"
)

func main() {
	fs := http.FileServer(http.Dir(staticPath))
	http.Handle(stringStatic, http.StripPrefix(stringStatic, fs))

	http.HandleFunc(rootPath, services.HtmlRead)
	http.HandleFunc(registerPath, services.Register)
	http.HandleFunc(getPath, services.Get)
	http.HandleFunc(updatePath, services.Update)
	http.HandleFunc(deletePath, services.Delete)

	log.Println(serverLink)
	err := http.ListenAndServe(localHost, nil)
	if err != nil {
		log.Println(err)
		return
	}
}
