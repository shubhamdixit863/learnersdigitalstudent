package processor

import (
	"fmt"
	"strings"
)

// WordCount counts total words across all files
func WordCount(linesCh <-chan string) {
	totalWords := 0

	for line := range linesCh {
		totalWords += len(strings.Fields(line))
	}

	fmt.Println("Total Word Count:", totalWords)
}
