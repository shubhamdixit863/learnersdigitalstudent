package services

import (
	"fmt"
	"sync"
	"time"
)

type LogEntry struct {
	Timestamp time.Time
	Level     string
	Message   string
}

var arr []LogEntry

func NewLog(t time.Time, lev string, mesg string) *LogEntry {

	return &LogEntry{
		Timestamp: time.Now(),
		Level:     lev,
		Message:   mesg,
	}
}
func Store(newch chan *LogEntry, wg *sync.WaitGroup) {
	defer wg.Done()
	val := <-newch
	arr = append(arr, *val)
	fmt.Println(arr)

}
