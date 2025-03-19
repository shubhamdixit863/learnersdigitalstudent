package services

import (
	"math/rand"
	"sync"
	"time"
)

func LogCreator(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	t := time.Now().Format("2006-01-02 15:04:05")
	slc := []string{" INFO: User logged in", " ERRR: Database connection failed", " WARN: Disk usage is high"}

	for i := 0; i < 8; i++ {
		r := rand.Intn(len(slc))
		ch <- t + slc[r]
	}

	close(ch)
}
