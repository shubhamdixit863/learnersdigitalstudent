package services

import (
	"sync"
)

var Mp = make(map[string]int)

func LogStorage(wg *sync.WaitGroup, logCh chan LogEntry) {
	defer wg.Done()

	for v := range logCh {
		Mp[v.Level]++
	}
}
