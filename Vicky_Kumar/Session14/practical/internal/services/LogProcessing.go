package services

import (
	"sync"
	"time"
)

type LogEntry struct {
	Timestamp time.Time
	Level     string
	Message   string
}

func NewLogEntry(tm time.Time, lvl string, msg string) *LogEntry {
	return &LogEntry{Timestamp: tm, Level: lvl, Message: msg}
}

func LogProcessing(wg *sync.WaitGroup, levelCh chan string, msgCh chan string, tmCh chan time.Time, logCh chan LogEntry) {
	defer wg.Done()

	var ok = true
	var lvl string
	lvl, ok = <-levelCh
	i := 1
	for ok {
		msg := <-msgCh
		tm := <-tmCh
		log := NewLogEntry(tm, lvl, msg)
		logCh <- *log
		i++
		lvl, ok = <-levelCh
	}
	close(logCh)
}
