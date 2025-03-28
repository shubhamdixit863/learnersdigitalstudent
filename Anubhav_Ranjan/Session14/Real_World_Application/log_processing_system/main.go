package main

import (
	"fmt"
	"log_processing_system/internal/services"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	rawLogChan := make(chan string, 15)
	processedLogChan := make(chan services.LogEntry, 15)
	doneChan := make(chan bool)

	var producerWg sync.WaitGroup
	var workerWg sync.WaitGroup

	logStorage := services.NewLogStorage(processedLogChan, doneChan)
	go logStorage.Store()

	numProducers := 3
	logsPerProducer := 10
	for i := 1; i <= numProducers; i++ {
		producerWg.Add(1)
		generator := services.NewLogGenerator(i, logsPerProducer, rawLogChan, &producerWg)
		go generator.Generate()
	}

	numWorkers := 5
	for i := 1; i <= numWorkers; i++ {
		workerWg.Add(1)
		processor := services.NewLogProcessor(i, rawLogChan, processedLogChan, &workerWg)
		go processor.Process()
	}

	go func() {
		producerWg.Wait()
		close(rawLogChan)
		fmt.Println("All producers finished, closed raw log channel")
	}()

	go func() {
		workerWg.Wait()
		close(processedLogChan)
		fmt.Println("All workers finished, closed processed log channel")
	}()

	dn := <-doneChan

	if dn {
		fmt.Println("Logging Processing Completed")
	}

	summary := logStorage.GetSummary()
	fmt.Printf("Log Level Count Summary: %v\n", summary)
}
