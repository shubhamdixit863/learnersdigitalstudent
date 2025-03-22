package services

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// ExtractWordsAndLinks fetches a webpage, extracts links and words
func ExtractWordsAndLinks(url string) ([]string, map[string]int) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to fetch:", url)
		return nil, nil
	}
	defer resp.Body.Close()

	links := []string{}
	wordCounts := make(map[string]int)
	tokenizer := html.NewTokenizer(resp.Body)

	for {
		tokenType := tokenizer.Next()
		switch tokenType {
		case html.ErrorToken:
			return links, wordCounts
		case html.StartTagToken:
			token := tokenizer.Token()
			if token.Data == "a" {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						links = append(links, attr.Val)
					}
				}
			}
		case html.TextToken:
			text := strings.ToLower(strings.TrimSpace(tokenizer.Token().Data))
			words := strings.Fields(text)
			for _, word := range words {
				wordCounts[word]++
			}
		}
	}
}
