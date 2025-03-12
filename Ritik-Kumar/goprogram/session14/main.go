package main

import (
	"fmt"
	"session14/services"
	"time"
)

func main() {
	logChannel := make(chan string)
	processedChannel := make(chan services.LogEntry)
	countChannel := make(chan map[string]int)

	numLogs := 20
	go services.GenerateLogs(logChannel, numLogs)

	go services.ProcessLogs(logChannel, processedChannel)
	//close(processedChannel)

	go func() {
		services.AggregateLogs(processedChannel, countChannel)
		close(countChannel)
	}()
	logCounts := <-countChannel
	fmt.Println("Log Level Count Summary:", logCounts)
	time.Sleep(2 * time.Second)
}
