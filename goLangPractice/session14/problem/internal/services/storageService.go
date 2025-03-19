package services

import (
	"sync"
)

func StoreLogs(logCount map[string]int, ch2 chan LogEntry, wg *sync.WaitGroup) {
	defer wg.Done()

	for entry := range ch2 {
		logCount[entry.Level]++
	}
}
