package processor

import (
	"fmt"
	"strings"
)

// FilterLines filters lines containing a keyword as an exact word match (case-insensitive)
func FilterLines(linesCh <-chan string, keyword string) {
	keyword = strings.ToLower(keyword) 

	for line := range linesCh {
		words := strings.Fields(strings.ToLower(line)) // Convert line to lowercase & split into words
		for _, word := range words {
			if word == keyword {
				fmt.Println("Filtered:", line)
				break
			}
		}
	}
}
