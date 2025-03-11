package services

import (
	"strings"
	"sync"
	"time"
)

type LogEntry struct {
	Timestamp time.Time
	Level     string
	Message   string
}

func LogProcessing(ch chan string, wg *sync.WaitGroup, logentrychannel chan LogEntry) {
	defer wg.Done()
	var entry LogEntry
	for v := range ch {
		str := strings.Split(v, ":")
		if str[0] == "INFO" {
			entry = LogEntry{
				Timestamp: time.Now(),
				Level:     str[0],
				Message:   str[1],
			}
		} else if str[0] == "ERROR" {
			entry = LogEntry{
				Timestamp: time.Now(),
				Level:     str[0],
				Message:   str[1],
			}

		} else if str[0] == "WARN" {
			entry = LogEntry{
				Timestamp: time.Now(),
				Level:     str[0],
				Message:   str[1],
			}
		}
		logentrychannel <- entry
	}
	close(logentrychannel)

}
