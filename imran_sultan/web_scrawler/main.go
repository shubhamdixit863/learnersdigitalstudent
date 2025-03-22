package main

import (
	"fmt"
	"web_scrawler/services"
)

func main() {
	startURL := "https://usf-cs272-s25.github.io/top10/Dracula%20%7C%20Project%20Gutenberg/index.html" // Replace with the website you want to crawl
	maxDepth := 2                                                                                      // Set crawling depth

	results := services.StartCrawling(startURL, maxDepth)

	// Print word-to-links mapping
	for word, links := range results {
		fmt.Printf("Word: %s\n", word)
		for link, count := range links {
			fmt.Printf("  %s (found %d times)\n", link, count)
		}
	}
}
