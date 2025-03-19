package main

import (
	"handsOn/internal/model"
	"handsOn/internal/services"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	logChan := make(chan string)
	processedChan := make(chan model.LogEntry)

	go services.GenerateService(logChan, &wg)
	go services.ParseLog(logChan, processedChan, &wg)
	go services.Aggregate(processedChan, &wg)

	wg.Add(3)

	wg.Wait()

}
