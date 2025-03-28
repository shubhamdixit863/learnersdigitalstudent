package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"assignment3/internal/services"
)


type Aggregator struct {
	filePath   string
	chunkChan  chan string
	resultChan chan map[string]int
}


func NewAggregator(filePath string) *Aggregator {
	return &Aggregator{
		filePath:   filePath,
		chunkChan:  make(chan string),
		resultChan: make(chan map[string]int),
	}
}


func (a *Aggregator) Run() {
	
	file, err := os.Open(a.filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	go services.Reader(file, a.chunkChan)           
	go services.FrequencyCounter(a.chunkChan, a.resultChan) 

	
	finalFreq := make(map[string]int)
	for freq := range a.resultChan {
		for word, count := range freq {
			finalFreq[word] += count
		}
	}

	fmt.Println("Word Frequencies:")
	for word, count := range finalFreq {
		fmt.Printf("%s: %d\n", word, count)
	}
}

func main() {
	
	filePath := filepath.Join("internal", "services", "words.txt")

	aggregator := NewAggregator(filePath)

	
	aggregator.Run()
}
