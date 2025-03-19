package services

import (
	"fmt"
	"sync"
	"time"
)

type LogEntry struct {
	Timestamp time.Time

	Level string

	Message string
}

func ProcessLogs(ch chan string, ch2 chan LogEntry, wg *sync.WaitGroup) {
	defer wg.Done()

	for logRec := range ch {
		fmt.Println(logRec)

		timestamp := logRec[:19]
		label := logRec[20:24]
		message := logRec[26:]

		parsedTime, err := time.Parse("2006-01-02 15:04:05", timestamp)
		if err != nil {
			fmt.Println("Error parsing timestamp:", err)
		}

		entry := LogEntry{
			Timestamp: parsedTime,
			Level:     label,
			Message:   message,
		}

		ch2 <- entry
	}
	close(ch2)
}
