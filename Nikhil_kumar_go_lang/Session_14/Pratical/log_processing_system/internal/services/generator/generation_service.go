package services

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func GenerateLogs(logChan chan string, logNums int, wg *sync.WaitGroup) {
	defer wg.Done()
	levels := []string{"INFO", "ERROR", "WARN"}
	msg := []string{" User logged in ", " Database connection failed ", " Disk usage is high "}
	// fmt.Println(levels)

	for i := 0; i < logNums; i++ {
		timestamp := time.Now()
		val := rand.Intn(len(levels))
		level := levels[val]
		message := msg[val]

		logEntry := fmt.Sprintf("[%s],%s,%s", timestamp, level, message)
		// fmt.Printf("%s - %s - %s\n", timestamp, level, logEntry)
		logChan <- logEntry
		// fmt.Println(logChan)
	}
	close(logChan)

}
