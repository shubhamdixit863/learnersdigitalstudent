package services

import (
	"log_processing_system/internal/services/model"
	"sync"
)

func StoreLogs(ProcessedLog chan model.LogEntry, storage chan map[string]int, wg *sync.WaitGroup) {
	defer wg.Done()
	logcount := make(map[string]int)
	for log := range ProcessedLog {
		logcount[log.Level]++
	}
	storage <- logcount
	// fmt.Println(storage)
	close(storage)
}
