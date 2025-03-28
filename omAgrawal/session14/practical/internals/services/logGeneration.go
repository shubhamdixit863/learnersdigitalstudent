package services

import (
	"math/rand"
	"sync"
	"time"
)

func LogGeneration(log chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	sl := []string{" INFO: User logged in", " ERRR: Database connection failed", "WARN: Disk usage is high"}

	for i := 0; i < 5; i++ {
		//fmt.Println("log generated")
		random := rand.Intn(3)
		log <- time.Now().String() + sl[random]
	}
	close(log)

}
