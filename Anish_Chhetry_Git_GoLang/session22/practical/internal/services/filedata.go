package services

import (
	"log"
	"net/http"
	"os"
)

const (
	fileLocation = "../files/"
)

func FileDataHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		filename := r.URL.Query().Get("filename")
		log.Println(filename)
		data, err := os.ReadFile(fileLocation + filename)
		if err != nil {
			log.Println("reading file error")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("reading file error"))
			return
		}
		log.Println(string(data))
		w.Write(data)
		return

	})

}
