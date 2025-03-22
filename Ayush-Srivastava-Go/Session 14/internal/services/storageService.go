package services

import "fmt"

type LogStorage struct {
	logCounts map[string]int
}

func NewLogStorage() *LogStorage {
	return &LogStorage{
		logCounts: map[string]int{
			"INFO":  0,
			"ERROR": 0,
			"WARN":  0,
		},
	}
}

func (ls *LogStorage) StoreLogs(processedChannel chan LogEntry) {
	for log := range processedChannel {
		ls.logCounts[log.Level]++
	}
}

func (ls *LogStorage) PrintSummary() {
	fmt.Println("Log Level Count Summary:", ls.logCounts)
}
