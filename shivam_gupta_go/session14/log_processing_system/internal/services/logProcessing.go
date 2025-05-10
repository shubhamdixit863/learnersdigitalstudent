package services

import (
	"strings"
)

type LogEntry struct {
	Timestamp string
	Level     string
	Message   string
}

func ProcessedLogs(logchannel chan string, ProcessedLog chan LogEntry) {
	for log := range logchannel {
		parts := strings.SplitN(log, ",",3)
    ProcessedLog <-LogEntry{Timestamp:parts[0],Level:parts[1],Message: parts[2] }
	}
	close(ProcessedLog)
}

