package services

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var Level = []string{"INFO", "ERROR", "WARN"}
var Message = []string{"User logged in", "Database connection failed", "Disk usage is high"}

func LogGenerator(wg *sync.WaitGroup, levelCh chan string, msgCh chan string, tmCh chan time.Time) {
	defer wg.Done()

	idx := rand.Intn(3)
	for i := 0; i < 15; i++ {
		levelCh <- Level[idx]
		msgCh <- Message[idx]
		tmCh <- time.Now()
		idx = rand.Intn(3)
		fmt.Printf("[%s] %s: %s\n", time.Now().Format("2006-01-02 15:04:05"), Level[idx], Message[idx])
	}
	close(levelCh)
	close(msgCh)
	close(tmCh)
}
