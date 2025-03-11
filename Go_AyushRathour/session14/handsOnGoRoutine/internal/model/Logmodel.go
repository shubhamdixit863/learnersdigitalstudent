package model

import "time"

type LogEntry struct {
	Timestamp time.Time

	Level string

	Message string
}
