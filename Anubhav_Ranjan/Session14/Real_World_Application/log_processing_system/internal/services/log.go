package services

import "time"

type LogEntry struct {
	Timestamp time.Time
	Level     string
	Message   string
}
