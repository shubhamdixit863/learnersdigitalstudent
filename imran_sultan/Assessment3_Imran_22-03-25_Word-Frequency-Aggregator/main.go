package main

import (
	"fmt"
	"handson3/processor"
	"handson3/reader"
	"handson3/utils"
	"sync"
)

const (
	chunkSize = 2
	fileName  = "book.txt"
)

func chunkToprocessor(chunks chan []string, results chan map[string]int, wg *sync.WaitGroup) {
	for chunk := range chunks {
		wg.Add(1)
		go processor.CountWords(chunk, results, wg)
	}

	wg.Wait()
	close(results)
}

func main() {
	chunks := make(chan []string)
	results := make(chan map[string]int)
	var wg sync.WaitGroup

	go reader.ReadChunks(fileName, chunkSize, chunks)

	go chunkToprocessor(chunks, results, &wg)
	wordCounts := utils.MergeResults(results)

	for word, count := range wordCounts {
		fmt.Printf("%s: %d\n", word, count)
	}
}
