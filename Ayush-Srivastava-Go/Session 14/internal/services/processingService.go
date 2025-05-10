package services

import (
	"fmt"
	"regexp"
	"time"
)

type LogEntry struct {
	Timestamp time.Time
	Level     string
	Message   string
}

type LogProcessor struct{}

func NewLogProcessor() *LogProcessor {
	return &LogProcessor{}
}

func (lp *LogProcessor) ProcessLogs(logChannel chan string, processedChannel chan LogEntry, numWorkers int) {
	for range numWorkers {
		go func() {
			for log := range logChannel {
				data, err := lp.parseLog(log)
				if err == nil {
					processedChannel <- data
				}
			}
		}()
	}
}

func (lp *LogProcessor) parseLog(log string) (LogEntry, error) {
	logPattern := `\[(.*?)\] (INFO|ERROR|WARN): (.*)`
	re := regexp.MustCompile(logPattern)
	matches := re.FindStringSubmatch(log)

	if len(matches) != 4 {
		return LogEntry{}, fmt.Errorf("invalid log format")
	}

	timestamp, err := time.Parse("2006-01-02 15:04:05", matches[1])
	if err != nil {
		return LogEntry{}, err
	}

	return LogEntry{
		Timestamp: timestamp,
		Level:     matches[2],
		Message:   matches[3],
	}, nil
}

