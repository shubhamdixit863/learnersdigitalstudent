package services

import (
	"strings"
	"sync"
)

func (cr *CrawlerService) IndexService(chCleaned chan string, chLink chan string, mu *sync.Mutex) {

	for text := range chCleaned {
		document := <-chLink
		words := strings.Fields(text)
		mu.Lock()
		for _, word := range words {
			lowerWord := strings.ToLower(strings.Trim(word, ".,!?;:'\"()[]{}<>"))

			if _, exists := cr.InvertedIndex[lowerWord]; !exists {
				cr.InvertedIndex[lowerWord] = make(map[string]int)
			}
			cr.InvertedIndex[lowerWord][document]++
		}
		mu.Unlock()
	}

}
