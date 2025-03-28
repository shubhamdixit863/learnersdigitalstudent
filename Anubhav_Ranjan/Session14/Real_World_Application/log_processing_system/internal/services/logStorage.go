package services

import (
	"fmt"
	"sync"
)

type LogStorage struct {
	inputChan  <-chan LogEntry
	logCounts  map[string]int
	countMutex sync.Mutex
	done       chan<- bool
}

func NewLogStorage(inputChan <-chan LogEntry, done chan<- bool) *LogStorage {
	return &LogStorage{
		inputChan: inputChan,
		logCounts: make(map[string]int),
		done:      done,
	}
}

func (ls *LogStorage) Store() {
	for entry := range ls.inputChan {
		ls.countMutex.Lock()
		ls.logCounts[entry.Level]++
		ls.countMutex.Unlock()
	}

	fmt.Println("Storage completed")
	ls.done <- true
}

func (ls *LogStorage) GetSummary() map[string]int {

	return ls.logCounts
}
