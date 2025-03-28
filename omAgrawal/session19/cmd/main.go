package main

import (
	"net/http"
	"session19/internals/handler"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("internals/directory")))
	http.HandleFunc("/register", handler.HandleRegister)
	http.HandleFunc("/get/user", handler.HandleGetAll)
	http.HandleFunc("/get/userid", handler.HandleGetById)
	http.HandleFunc("/update", handler.HandleUpdate)
	http.HandleFunc("/delete", handler.HandleDelete)

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		return
	}

}
