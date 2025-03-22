package processor

import (
	"strings"
	"sync"
)

func CountWords(lines []string, out chan<- map[string]int, wg *sync.WaitGroup) {
	defer wg.Done()

	wordFreq := make(map[string]int)
	for _, line := range lines {
		words := strings.Fields(strings.ToLower(line))
		for _, word := range words {
			word = strings.Trim(word, ".,!?;:\"'()[]{}")
			wordFreq[word]++
		}
	}
	out <- wordFreq
}
