package services

import (
	"fmt"
	"handsOn/internal/model"
	"sync"
)

var (
	levelCounts = make(map[string]int)
	//aggWg       sync.WaitGroup
)

func Aggregate(ProcessedChan chan model.LogEntry, aggWg *sync.WaitGroup) {
	defer aggWg.Done()
	for entry := range ProcessedChan {
		levelCounts[entry.Level]++

	}
	fmt.Println("Log level Count summary: ", levelCounts)
	//close(ProcessedChan)
}
