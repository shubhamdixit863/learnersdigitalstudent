package services

import (
	"regexp"
	"time"
)

type LogEntry struct {
	Timestamp time.Time
	Level     string
	Message   string
}

func ProcessLogs(logChannel <-chan string, processedChannel chan<- LogEntry) {
	re := regexp.MustCompile(`\[(.*?)\] (\w+): (.*)`)

	for log := range logChannel {
		matches := re.FindStringSubmatch(log)
		if len(matches) == 4 {
			parsedTime, _ := time.Parse("2006-01-02 15:04:05", matches[1])
			entry := LogEntry{
				Timestamp: parsedTime,
				Level:     matches[2],
				Message:   matches[3],
			}
			processedChannel <- entry
		}
	}
	close(processedChannel)
}
