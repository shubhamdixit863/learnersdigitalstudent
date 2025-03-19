package main

import (
	"fmt"
	"session14/problem/internal/services"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan string)
	ch2 := make(chan services.LogEntry)
	logCount := make(map[string]int)

	go services.LogCreator(ch, &wg)
	go services.ProcessLogs(ch, ch2, &wg)
	go services.StoreLogs(logCount, ch2, &wg)

	wg.Add(3)
	wg.Wait()

	fmt.Println("Log Level Count Summary: ", logCount)
}
