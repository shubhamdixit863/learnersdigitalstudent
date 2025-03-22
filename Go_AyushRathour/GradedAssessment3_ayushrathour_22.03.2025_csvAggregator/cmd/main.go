package main

import (
	"assessment3/internal/processor"
	"assessment3/internal/reader"
	"assessment3/internal/types"
	"assessment3/internal/utils"
	"fmt"
	"log"
	"sync"
)

const (
	csvFilePath = "C:/Users/Asus/Desktop/GradedAssessment3/internal/data/transactions.csv"
	chunkSize   = 2
)

func main() {
	chunks, err := reader.ReadCSV(csvFilePath, chunkSize)
	if err != nil {
		log.Fatalf("Failed to read CSV file: %v", err)
	}

	userSpend := make(types.UserSpend)
	ch := make(chan types.UserSpend, len(chunks))
	var wg sync.WaitGroup
	var mu sync.Mutex

	processor := processor.NewChunkProcessor() // Dependency Injection

	for _, chunk := range chunks {
		wg.Add(1)
		go func(chunk [][2]string) {
			defer wg.Done()
			ch <- processor.Process(chunk)
		}(chunk)
	}

	// Close the channel once all goroutines are done
	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		utils.MergeMaps(userSpend, result, &mu)
	}

	fmt.Println("Total Spend Per User:")
	for user, spend := range userSpend {
		fmt.Printf("%s: %.2f\n", user, spend)
	}
}
