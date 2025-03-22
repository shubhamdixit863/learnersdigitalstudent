package main

import (
	"logSys/internal/services"
	"sync"
	"time"
)

func main() {
	numLogs := 20      
	numWorkers := 3    

	logChannel := make(chan string, numLogs)
	processedChannel := make(chan services.LogEntry, numLogs)

	logGenerator := services.NewLogGenerator()
	logProcessor := services.NewLogProcessor()
	logStorage := services.NewLogStorage()
  
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		logProcessor.ProcessLogs(logChannel, processedChannel, numWorkers)
	}()

	go func() {
		defer wg.Done()
		logStorage.StoreLogs(processedChannel)
	}()

	logGenerator.GenerateLogs(logChannel, numLogs)
	close(logChannel)

	logGenerator.PrintGeneratedLogs()

	time.Sleep(time.Second)
	close(processedChannel)

	wg.Wait()

	logStorage.PrintSummary()
}
