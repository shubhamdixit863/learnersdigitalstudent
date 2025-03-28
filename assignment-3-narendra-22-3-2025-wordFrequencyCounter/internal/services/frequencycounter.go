package services

import (
	"fmt"
)


func FrequencyCounter(chunkChan <-chan string, resultChan chan<- map[string]int) {
	freq := make(map[string]int)

	for line := range chunkChan {
		words := Chunker(line) 
		for _, word := range words {
			freq[word]++
		}
	}

	
	if len(freq) == 0 {
		fmt.Println("Warning: No words found in file!")
	}

	resultChan <- freq
	close(resultChan)
}















