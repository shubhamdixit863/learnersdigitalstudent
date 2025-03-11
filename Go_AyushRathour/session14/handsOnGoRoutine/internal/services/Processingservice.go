package services

import (
	"handsOn/internal/model"
	"strings"
	"sync"
	"time"
)

func ParseLog(logStr chan string, ProcessedChan chan model.LogEntry, processLog *sync.WaitGroup) {
	defer processLog.Done()
	for log := range logStr {
		parts := strings.SplitN(log, "]", 2)
		timeStampStr := parts[0][1:]
		timestamp, _ := time.Parse("2006-01-02 15:04:05", timeStampStr)
		levelMsg := strings.SplitN(parts[1], ":", 2)

		newData := model.LogEntry{
			Timestamp: timestamp,
			Level:     levelMsg[0],
			Message:   levelMsg[1],
		}

		ProcessedChan <- newData
	}

	close(ProcessedChan)
}
