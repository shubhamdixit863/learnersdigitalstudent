package main

import (
	"fmt"
	"log"
	"net/http"
	"url_shortner_application/internal/services"
)

func handleRedirect(shortener *services.URLShortener) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shortCode := r.URL.Path[1:]

		originalURL, err := shortener.GetOriginalURL(shortCode)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Redirect(w, r, originalURL, http.StatusFound)
	}
}

func main() {
	shortener := services.NewURLShortener()

	url := "https://learnersdigital.com"

	shortCode, err := shortener.ShortenURL(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Short URL: http://localhost:8080/%s\n", shortCode)

	http.HandleFunc("/", handleRedirect(shortener))

	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
