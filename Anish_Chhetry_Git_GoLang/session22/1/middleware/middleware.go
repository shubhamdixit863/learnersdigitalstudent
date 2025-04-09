package middleware

import (
	"fmt"
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//print out the request

		fmt.Println(r.URL.Path)

		fmt.Println(r.URL.Query().Get("token"))
		//now you have access to request and response
		if r.URL.Query().Get("token") == "" {
			//you can directly send a response
			w.Write([]byte("Token required"))
			return
		}
		//this should be called when everything is fine
		next.ServeHTTP(w, r)
	})
}

//use net http package
//you have to create two handlers which gives the file list and
//another gives the file data as per the name passed
//you have to create a log middleware that logs following things
//first logs all things about request in structure format, agent, header
//you have to create a map with a user list and the files they can access
//if the user tries to access any other file other than what they have access to send 405 status forbidden
