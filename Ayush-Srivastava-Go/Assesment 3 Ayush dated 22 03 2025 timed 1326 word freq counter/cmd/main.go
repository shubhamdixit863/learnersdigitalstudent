package main

import (
	"filemonitor/internal/monitor"
	"filemonitor/internal/utils"
	"fmt"
	"sync"
)

func main() {

	fileChan := make(chan string, len(utils.Files))

	var wg sync.WaitGroup

	mon := monitor.NewMonitor(fileChan, &wg)
	mon.StartProcessing()

	for _, file := range utils.Files {
		fileChan <- file
	}

	close(fileChan)
	wg.Wait()

	fmt.Println("All files processed")
}