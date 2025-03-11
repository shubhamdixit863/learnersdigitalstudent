package services

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func GenerateService(logChan chan string, producerWg *sync.WaitGroup) {
	//var producerWg sync.WaitGroup
	defer producerWg.Done()
	levels := []string{"INFO", "WARN", "ERROR"}
	messages := []string{"user logged in", "Database connection failed", "Disk usage is high"}
	for j := 0; j < 5; j++ {
		idx := rand.Intn(3)
		timestamp := time.Now().Add(time.Second * time.Duration(j)).Format("2006-01-02 15:04:05")
		level := fmt.Sprintf(levels[idx])
		message := fmt.Sprintf(messages[idx])

		logStr := fmt.Sprintf("[%s] %s: %s", timestamp, level, message)
		logChan <- logStr

	}
	close(logChan)
}
