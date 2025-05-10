package services

import (
	"sync"
)

func LogGeneration(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 8; i++ {
		log := "INFO: User logged in"
		ch <- log
	}
	for i := 0; i < 5; i++ {
		log := "ERROR: Database connection failed"
		ch <- log
	}
	for i := 0; i < 7; i++ {
		log := "WARN:  Disk usage is high"
		ch <- log
	}
	close(ch)
}
