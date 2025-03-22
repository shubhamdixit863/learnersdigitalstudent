package main

import (
	"fmt"
	"log"
	"sync"
	"wordfrequencycounter/internal/services"
	"wordfrequencycounter/internal/models"
)

const filereaderror = "File read error: %v"
const workererror ="Worker %d encountered an error: %v"

func main() {
	ch := make(chan []byte)
	resultCh := make(chan map[string]int)
	doneCh := make(chan bool)

	
	var wg sync.WaitGroup

	
	go func() {
		
		if err := services.ReadFile(models.FilePath, models.ChunkSize, ch); err != nil {
			log.Fatalf(filereaderror, err)
		}
		
	}()

	
	for i := 0; i < models.NumWorkers; i++ {
		wg.Add(1) 
		go func(workerID int) {
			defer wg.Done() 
			for data := range ch {
				words, err := services.WordScanner(data)
				if err != nil {
					log.Printf(workererror, workerID, err)
					continue
				}
				freqMap := services.CalculateFrequency(words)
				resultCh <- freqMap
			}
		}(i)
	}

	
	 go func() {
		wg.Wait() 
		close(resultCh) 
	}() 

	
	finalFreq := make(map[string]int)
	go services.AggregateResults(resultCh, finalFreq, doneCh)

	
	<-doneCh

	
	fmt.Println("Top Word Frequencies:")
	count := 0
	for word, freq := range finalFreq {
		fmt.Printf("%s: %d\n", word, freq)
		count++
		if count == models.TopWordCount {
			break
		}
	}
}
