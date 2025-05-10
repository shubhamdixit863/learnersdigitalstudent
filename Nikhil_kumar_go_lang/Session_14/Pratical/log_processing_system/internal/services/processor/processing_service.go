package services

import (
	"log_processing_system/internal/services/model"
	"strings"
	"sync"
)

func ProcessingLogs(logChan chan string, ProcessedLog chan model.LogEntry, wg *sync.WaitGroup) {
	defer wg.Done()
	for l_entry := range logChan {
		str_split := strings.Split(l_entry, ",")
		ProcessedLog <- model.LogEntry{Timestamp: str_split[0], Level: str_split[1], Message: str_split[2]}
        
	}
	close(ProcessedLog)

}
