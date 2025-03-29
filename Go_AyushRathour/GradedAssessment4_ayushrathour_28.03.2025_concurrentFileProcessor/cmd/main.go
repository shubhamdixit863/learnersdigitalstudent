package main

import (
	"concurrentFileProcessor/internal/fileReader"
	"concurrentFileProcessor/internal/processor"
	"fmt"
	"sync"
)

const (
	dirPath    = "./sampleTexts"
	modeFilter = "line-filter"
	modeCount  = "word-count"
	modeAPI    = "api-call"
	modeRetry  = "retryable-api"
	filterKey  = "Error"
)

func main() {
	modes := []string{modeFilter, modeCount, modeAPI, modeRetry}

	for _, mode := range modes {
		lineChan := make(chan string, 100)
		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			defer wg.Done()
			fileReader.ReadFiles(dirPath, lineChan)
			close(lineChan)
		}()

		keyWord := ""
		if mode == modeFilter {
			keyWord = filterKey
		}

		result := processor.ProcessLines(lineChan, mode, keyWord)
		wg.Wait()
		fmt.Printf("Mode: %s, Result: %v\n", mode, result)
	}
}
