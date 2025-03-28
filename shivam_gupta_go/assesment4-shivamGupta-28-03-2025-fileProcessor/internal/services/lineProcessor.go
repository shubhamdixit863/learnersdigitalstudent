package services

import (
	"fileProcessor/internal/models"
	"fmt"
	"strings"
	"sync"
)

const (
 s1 = "Filtered: %s:%d - %s\n"
 s2 = "Words: %s:%d - %d words\n"
 s3 = "API: %s:%d - Success: %v\n"
 s4 = "Retry API: %s:%d - Success: %v\n"
 err31 = "error"
)


func ProcessLines(lines <-chan models.LineWithMetadata) {
	var wg sync.WaitGroup

	for line := range lines {
		wg.Add(1)
		go func(l models.LineWithMetadata) {
			defer wg.Done()


			filterResult := FilterLine(l,err31 )
			wordCount := WordCount(l)
			apiResult := APICall(l)
			retryResult := RetryableAPICall(l)

			
			if filterResult.Keep {
				fmt.Printf(s1, l.FileName, l.LineNum, l.Line)
			}
			fmt.Printf(s2, l.FileName, l.LineNum, wordCount.WordCount)
			fmt.Printf(s3, l.FileName, l.LineNum, apiResult.APISuccess)
			fmt.Printf(s4, l.FileName, l.LineNum, retryResult.APISuccess)
		}(line)
	}

	wg.Wait()
}

func FilterLine(line models.LineWithMetadata, keyword string) models.ProcessResult {
	return models.ProcessResult{
		FileName: line.FileName,
		LineNum:  line.LineNum,
		Line:     line.Line,
		Keep:     strings.Contains(line.Line, keyword),
	}
}

func WordCount(line models.LineWithMetadata) models.ProcessResult {
	return models.ProcessResult{
		FileName:  line.FileName,
		LineNum:   line.LineNum,
		Line:      line.Line,
		WordCount: len(strings.Fields(line.Line)),
	}
}
