package middleware

import (
	"log"
	"net/http"
)

var UserAccess = map[string][]string{
	"user1": {"a.txt", "c.txt"},
	"user2": {"b.txt", "c.txt"},
}

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Agent: ", r.UserAgent())
		log.Println("Header: ", r.Header)
		next.ServeHTTP(w, r)
	})
}

func UserAccessMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := r.URL.Query().Get("username")
		filename := r.URL.Query().Get("filename")

		if username == "" || filename == "" {
			log.Println("username or filename is empty")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("username or filename is empty"))
			return
		}
		for _, allowedFiles := range UserAccess[username] {
			if filename == allowedFiles {
				next.ServeHTTP(w, r)
				return
			}
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(filename + " filename is not allowed"))
		return
	})
}
