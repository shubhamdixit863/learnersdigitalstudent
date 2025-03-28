package processor

import (
	"concurrent-file-processor/constants"
	"fmt"
)

// ProcessLines routes processing based on mode
func ProcessLines(linesCh <-chan string, mode, keyword string) {
	switch mode {
	case "filter":
		FilterLines(linesCh, keyword)
	case "wordcount":
		WordCount(linesCh)
	case "apicall":
		APICall(linesCh)
	case "retryapi":
		RetryableAPICall(linesCh)
	default:
		fmt.Println(constants.InvalidMode)
	}
}
