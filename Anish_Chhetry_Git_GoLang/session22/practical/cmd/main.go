package main

import (
	"fmt"
	"net/http"
	"practical/internal/middleware"
	"practical/internal/services"
)

func main() {
	http.Handle("/filenames", middleware.LogMiddleware(services.FileNamesHandler()))
	http.Handle("/file", middleware.LogMiddleware(middleware.UserAccessMiddleware(services.FileDataHandler())))
	fmt.Println("To check the files in the directory:\nhttp://localhost:8080/filenames")
	fmt.Println("To enter username and book name:\nhttp://localhost:8080/file?username=user1&filename=a.txt")
	http.ListenAndServe(":8080", nil)
}
