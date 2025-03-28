package main

import (
	"fmt"
	"session12/session14/practical/internals/services"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	logg := make(chan string)
	stor := make(chan *services.LogEntry)
	go services.LogGeneration(logg, &wg)
	go services.LogProcessing(logg, stor, &wg)
	go services.Store(stor, &wg)
	wg.Add(3)
	wg.Wait()
	fmt.Println(services.Mapp)
}
