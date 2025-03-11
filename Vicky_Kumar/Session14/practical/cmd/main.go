package main

import (
	"fmt"
	"practical/internal/services"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	level := make(chan string)
	msg := make(chan string)
	tm := make(chan time.Time)
	log := make(chan services.LogEntry)

	wg.Add(3)
	go services.LogGenerator(wg, level, msg, tm)
	go services.LogProcessing(wg, level, msg, tm, log)
	go services.LogStorage(wg, log)
	wg.Wait()

	fmt.Println("Log Level Count Summary:", services.Mp)
}
