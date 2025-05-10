package main

import (
	"log"
	"os"
	"sort"
	"sync"
	"wordFrequencyAggregator/internal/services"
	"wordFrequencyAggregator/internal/utils"
	wordprocessor "wordFrequencyAggregator/internal/wordProcessor"
)

func main() {
	if len(os.Args) != 2 {
		log.Println(utils.FILE_ERR)
		return
	}

	filePath := os.Args[1]
	chunkSize := 1000

	chunks, err := services.ReadFileInChunks(filePath, chunkSize)
	if err != nil {
		log.Printf(utils.READ_ERR, err)
		return
	}

	results := make(chan services.WordFrequency, len(chunks))
	errChan := make(chan error, len(chunks))
	var wg sync.WaitGroup
	var mutex sync.Mutex
	processor := wordprocessor.WordProcessor{}
	for _, chunk := range chunks {
		wg.Add(1)
		go wordprocessor.ChunkProcessor(chunk, processor, results, &wg, errChan)
	}

	go func() {
		wg.Wait()
		close(results)
		close(errChan)
	}()

	var frequencies []services.WordFrequency
	var errors []error

	for freq := range results {
		frequencies = append(frequencies, freq)
	}

	for err := range errChan {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		log.Println(utils.ERR)
		for _, err := range errors {
			log.Println(err)
		}
		return
	}

	mergedFrequency := services.MergeWordFrequency(frequencies, &mutex)

	type wordCount struct {
		word  string
		count int
	}
	var wordCounts []wordCount
	for word, count := range mergedFrequency {
		wordCounts = append(wordCounts, wordCount{word, count})
	}

	sort.Slice(wordCounts, func(i, j int) bool {
		return wordCounts[i].count > wordCounts[j].count
	})

	log.Println(utils.WORD_FRE_RES)
	log.Println(utils.DOT)
	for _, wc := range wordCounts {
		log.Printf(utils.WORD_COUNT, wc.word, wc.count)
	}

	log.Printf(utils.WORDCOUNT, len(wordCounts))
}
