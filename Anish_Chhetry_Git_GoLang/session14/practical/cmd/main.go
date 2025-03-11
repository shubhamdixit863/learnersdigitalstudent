package main

import (
	"fmt"
	"practical/internal/services"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	logchannel := make(chan string)
	logentrychannel := make(chan services.LogEntry)
	go services.LogProcessing(logchannel, &wg, logentrychannel)
	go services.LogGeneration(logchannel, &wg)
	go services.StorageService(logentrychannel, &wg)
	wg.Add(3)
	wg.Wait()
	fmt.Println("Log Level Count Summary:", services.FinalMap)

}
