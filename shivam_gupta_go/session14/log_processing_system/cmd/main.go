package main

import (
	"fmt"
	"session14/log_processing_system/internal/services"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	logchannel := make(chan string)
	processedLog := make(chan services.LogEntry)
	resultLog := make(chan map[string]int)

  go func(){ wg.Done()
			 services.ProcessedLogs(logchannel, processedLog)}()
 
	go  func(){ wg.Done()
		 services.StoreLogs(processedLog, resultLog)}()

	go func (){wg.Done()
		 services.GenLogs(logchannel, 10)}()

	logCounts := <-resultLog
	wg.Wait()
	fmt.Println("Log Counts by Level:", logCounts)
}
