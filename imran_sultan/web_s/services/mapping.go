package services

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"web_s/storage"
)

func GetPlainTextFromURL(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	htmlContent := string(body)

	// Remove HTML tags using regex
	re := regexp.MustCompile("<[^>]*>")
	plainText := re.ReplaceAllString(htmlContent, " ")

	// Trim and normalize spaces
	plainText = strings.Join(strings.Fields(plainText), " ")

	return plainText, nil
}

func IterateText(text string, url string) {
	words := strings.Split(text, " ")
	for _, word := range words {
		// Initialize the inner map if the word is not present
		if _, exists := storage.Mp[word]; !exists {
			storage.Mp[word] = make(map[string]int)
		}

		// Increment the count for the given URL
		storage.Mp[word][url]++
	}
}

func WordTolink(url string) {
	text, err := GetPlainTextFromURL(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	IterateText(text, url)
}
