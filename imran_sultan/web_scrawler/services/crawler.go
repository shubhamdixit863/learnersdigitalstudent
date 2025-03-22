package services

import (
	"fmt"
	"sync"
	storage "web_scrawler/storege"
)

// StartCrawling initializes the web crawler
func StartCrawling(startURL string, maxDepth int) map[string]map[string]int {
	wordMap := storage.NewStorage()
	visited := make(map[string]bool)
	var wg sync.WaitGroup

	var crawl func(url string, depth int)
	crawl = func(url string, depth int) {
		defer wg.Done()
		if depth > maxDepth || visited[url] {
			return
		}
		visited[url] = true

		fmt.Println("Crawling:", url)

		// Extract words and links
		links, foundWords := ExtractWordsAndLinks(url)

		// Store word occurrences
		for word, count := range foundWords {
			wordMap.Add(word, url, count)
		}

		// Crawl discovered links
		for _, link := range links {
			if !visited[link] {
				wg.Add(1)
				go crawl(link, depth+1)
			}
		}
	}

	wg.Add(1)
	go crawl(startURL, 0)
	wg.Wait()

	return wordMap.GetData()
}
