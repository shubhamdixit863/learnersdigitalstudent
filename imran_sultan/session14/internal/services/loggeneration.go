package services

import (
	"math/rand"
	"sync"
	"time"
)

func LogGeneration(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	t := time.Now()
	Time := t.String()
	sl := []string{"INFO: User logged in", "ERROR: Database connection failed", "WARN: Disk usage is high"}
	indx := rand.Intn(3)
	str := ""
	str = str + Time + sl[indx]
	ch <- str
	close(ch)

}
