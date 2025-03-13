package main

import (
	"fmt"
	generator "log_processing_system/internal/services/generator"
	log_storage "log_processing_system/internal/services/log_storage"
	"log_processing_system/internal/services/model"
	processor "log_processing_system/internal/services/processor"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	
	wg.Add(3)

	logChan := make(chan string)
	processed := make(chan model.LogEntry)
	storage := make(chan map[string]int)

	go generator.GenerateLogs(logChan, 10, &wg)
	go processor.ProcessingLogs(logChan, processed, &wg)
	go log_storage.StoreLogs(processed, storage, &wg)

	logCounts := <-storage

	wg.Wait()

	fmt.Println(logCounts)
}
