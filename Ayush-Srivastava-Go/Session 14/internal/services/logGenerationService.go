package services

import (
	"fmt"
	"math/rand"
	"time"
)

type LogGenerator struct {
	logs map[string][]string
}

func NewLogGenerator() *LogGenerator {
	return &LogGenerator{
		logs: map[string][]string{
			"INFO": {
				"User logged in",
				"File uploaded successfully",
				"Settings updated",
			},
			"ERROR": {
				"Database connection failed",
				"Failed to load configuration",
				"Unexpected server crash",
			},
			"WARN": {
				"Disk usage is high",
				"Memory consumption nearing limit",
				"Unusual login attempt detected",
			},
		},
	}
}

func (lg *LogGenerator) GenerateLogs(logChannel chan string, numLogs int) {
	levels := []string{"INFO", "ERROR", "WARN"}

	for range numLogs {
		level := levels[rand.Intn(len(levels))]
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		message := lg.logs[level][rand.Intn(len(lg.logs[level]))]

		formattedLog := fmt.Sprintf("[%s] %s: %s", timestamp, level, message)
		logChannel <- formattedLog
	}
}


func (lg *LogGenerator) PrintGeneratedLogs() {
	for level, messages := range lg.logs {
		for _, message := range messages {
			fmt.Printf("[%s] %s: %s\n", time.Now().Format("2006-01-02 15:04:05"), level, message)
		}
	}
}