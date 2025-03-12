package services

import (
	"fmt"
	"math/rand"
	"time"
)

var logLevels = []string{"INFO", "ERROR", "WARN"}

func GenerateLogs(logChannel chan<- string, numLogs int) {
	for i := 0; i < numLogs; i++ {
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		level := logLevels[rand.Intn(len(logLevels))]
		message := fmt.Sprintf("[%s] %s: log %d", timestamp, level, i+1)
		logChannel <- message
		time.Sleep(500 * time.Millisecond) 
	}
	close(logChannel)
}
