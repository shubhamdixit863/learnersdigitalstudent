package services

import (
	"fmt"
	"time"
)

type LogEntry struct {
	Timestamp time.Time
	Level     string
	Message   string
}

var Loglist = make([]*LogEntry, 3)

var Mapp = make(map[string]int)

func NewLog(tm string, lev string, loggg string) *LogEntry {
	const timeFormat = "2006-01-02 15:04:05.0000000"
	fmt.Println("time", tm)
	tme, err := time.Parse(timeFormat, tm)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return nil
	}
	return &LogEntry{Timestamp: tme, Level: lev, Message: loggg}
}
