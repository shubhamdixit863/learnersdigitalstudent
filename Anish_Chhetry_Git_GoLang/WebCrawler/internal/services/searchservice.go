package services

import (
	"fmt"
	"strings"
	"sync"
)

func (cr *CrawlerService) Search(searchFor string, mu *sync.Mutex) ([]string, error) {
	mu.Lock()
	defer mu.Unlock()
	searchFor = strings.ToLower(searchFor)
	urlSlice := []string{}

	if urls, exists := cr.InvertedIndex[searchFor]; exists {
		fmt.Printf("Results for \"%s\":\n", searchFor)
		for url, count := range urls {
			urlSlice = append(urlSlice, fmt.Sprintf("%s \n(found %d times).\n", url, count))
		}
		return urlSlice, nil
	}
	return nil, fmt.Errorf("No results found for '%s'.\n", searchFor)
}
