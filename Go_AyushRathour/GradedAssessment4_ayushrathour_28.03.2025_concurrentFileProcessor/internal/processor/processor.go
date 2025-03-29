package processor

import (
	"concurrentFileProcessor/internal/apiHandler"
	"fmt"
	"strings"
)

func ProcessLines(lineChan <-chan string, mode, keyword string) interface{} {
	switch mode {
	case "line-filter":
		return filterLines(lineChan, keyword)
	case "word-count":
		return countWords(lineChan)
	case "api-call":
		return apiHandler.CallAPI(lineChan)
	case "retryable-api":
		return apiHandler.CallAPIWithRetry(lineChan)
	default:
		fmt.Println("Invalid mode")
		return nil

	}
}

func filterLines(lineChan <-chan string, keyword string) []string {
	var result []string
	for line := range lineChan {
		if strings.Contains(line, keyword) { // Filters lines containing a specific keyword
			result = append(result, line)
		}
	}
	return result
}

func countWords(lineChan <-chan string) int {
	wordCount := 0
	for line := range lineChan { // Counts total words in all lines
		wordCount += len(strings.Fields(line))
	}
	return wordCount
}
