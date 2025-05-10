package services

import (
	"log"
	"net/http"
	"os"
)

func FileNamesHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		files, err := os.ReadDir(fileLocation)
		if err != nil {
			log.Println("could not read directory")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("could not read directory"))
			return
		}
		w.Write([]byte("File Names in the directory:"))
		for _, file := range files {
			w.Write([]byte("\n" + file.Name()))
		}

	})

}
