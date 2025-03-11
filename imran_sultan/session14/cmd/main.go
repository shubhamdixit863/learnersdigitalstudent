package main

import (
	"session14/internal/services"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan string)
	newch := make(chan *services.LogEntry)
	go services.LogGeneration(ch, &wg)
	go services.LogProcessing(ch, newch, &wg)
	go services.Store(newch, &wg)
	wg.Add(3)
	wg.Wait()

}
