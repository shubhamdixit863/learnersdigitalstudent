package main

import (
	"concurrentfileprocessor/internal/reader"
	"concurrentfileprocessor/internal/services"
	"flag"
	"fmt"
	"sync"
)
const (
	ModeFilter           = "filter"
	ModeWordCount        = "wordcount"
	ModeAPICall          = "apicall"
	ModeRetryableAPICall = "retryableapicall"
	directorymessage            = "Directory containing .txt files"
	emptystring =""
	directory = "dir"
	mode = "mode"
	filter = "filter"
	modemessage ="Processing mode: filter | wordcount | apicall | retryableapicall"
	keyword ="keyword"
	error ="error"
	keywordmessage= "Keyword to filter lines (only for 'filter' mode)"
	validdirectory ="Please provide a directory using -dir"
	invalidmode ="Invalid mode. Use filter, wordcount, apicall, or retryableapicall"
)


func main() {
	
	dirPath := flag.String(directory, emptystring, directorymessage)
	mode := flag.String(mode, filter, modemessage)
	keyword := flag.String(keyword, error, keywordmessage)
	flag.Parse()

	if *dirPath == emptystring {
		fmt.Println(validdirectory)
		return
	}

	linesChan := make(chan string, 100)
	var wg sync.WaitGroup

	
	go func() {
		reader.ReadFiles(*dirPath, linesChan, &wg)
		wg.Wait()
		close(linesChan) 
	}()

	// Process lines based on mode
	switch *mode {
	case ModeFilter:
		services.FilterLines(linesChan, *keyword)
	case ModeWordCount:
		services.CountWords(linesChan)
	case ModeAPICall:
		services.SendToAPI(linesChan)
	case ModeRetryableAPICall:
		services.SendToAPIWithRetry(linesChan)
	default:
		fmt.Println(invalidmode)
	}
}
