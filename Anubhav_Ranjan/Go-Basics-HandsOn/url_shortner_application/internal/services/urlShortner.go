package services

import (
	"fmt"
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const codeSize = 8

type URLShortener struct {
	urls map[string]URL
}

func NewURLShortener() *URLShortener {
	shortener := &URLShortener{
		urls: make(map[string]URL),
	}

	go shortener.cleanupExpiredURLs()

	return shortener
}

func generateShortCode() string {
	shortCode := make([]byte, codeSize)
	for i := range shortCode {
		shortCode[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortCode)
}

func (s *URLShortener) ShortenURL(originalURL string) (string, error) {
	if originalURL == "" {
		return "", fmt.Errorf("URL cannot be empty")
	}

	var shortCode string

	for {
		shortCode = generateShortCode()

		if _, exists := s.urls[shortCode]; !exists {
			break
		}
	}

	s.urls[shortCode] = URL{
		OriginalURL:  originalURL,
		CreationTime: time.Now(),
	}

	return shortCode, nil
}

func (s *URLShortener) GetOriginalURL(shortCode string) (string, error) {
	urlData, exists := s.urls[shortCode]
	if !exists {
		return "", fmt.Errorf("short code not found")
	}

	if time.Since(urlData.CreationTime) > 1*time.Hour {
		return "", fmt.Errorf("URL has expired")
	}

	return urlData.OriginalURL, nil
}

func (s *URLShortener) cleanupExpiredURLs() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		for code, urlData := range s.urls {
			if time.Since(urlData.CreationTime) > 1*time.Hour {
				delete(s.urls, code)
			}
		}

	}
}
