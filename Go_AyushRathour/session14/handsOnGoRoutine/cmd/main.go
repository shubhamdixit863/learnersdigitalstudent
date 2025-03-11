package main

import (
	"handsOn/internal/model"
	"handsOn/internal/services"
	"sync"
)

func main() {

	var generatorLog sync.WaitGroup
	var processLog sync.WaitGroup
	var aggregatorLog sync.WaitGroup

	logChan := make(chan string)
	processedChan := make(chan model.LogEntry)

	go services.GenerateService(logChan, &generatorLog)
	go services.ParseLog(logChan, processedChan, &processLog)
	go services.Aggregate(processedChan, &aggregatorLog)

	generatorLog.Add(1)
	processLog.Add(1)
	aggregatorLog.Add(1)

	generatorLog.Wait()
	processLog.Wait()
	aggregatorLog.Wait()

}
