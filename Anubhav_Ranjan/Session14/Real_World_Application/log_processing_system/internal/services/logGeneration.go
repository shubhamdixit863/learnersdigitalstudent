package services

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type LogGenerator struct {
	id         int
	count      int
	outputChan chan<- string
	wg         *sync.WaitGroup
	levels     []string
	messages   []string
}

func NewLogGenerator(id int, count int, outputChan chan<- string, wg *sync.WaitGroup) *LogGenerator {
	return &LogGenerator{
		id:         id,
		count:      count,
		outputChan: outputChan,
		wg:         wg,
		levels:     []string{"INFO", "ERROR", "WARN"},
		messages: []string{
			"User logged in",
			"Database connection failed",
			"Disk usage is high",
			"API request processed",
			"Cache miss",
			"File not found",
			"Authentication successful",
			"Session expired",
			"Memory usage spike",
		},
	}
}

func (lg *LogGenerator) Generate() {
	defer lg.wg.Done()

	for i := 0; i < lg.count; i++ {
		level := lg.levels[rand.Intn(len(lg.levels))]
		message := lg.messages[rand.Intn(len(lg.messages))]
		timestamp := time.Now().Format("2006-01-02 15:04:05")

		logEntry := fmt.Sprintf("[%s] %s: %s", timestamp, level, message)

		lg.outputChan <- logEntry

		time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
	}
	fmt.Printf("Producer %d completed\n", lg.id)
}
