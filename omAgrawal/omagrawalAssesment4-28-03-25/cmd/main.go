package main

import (
	"omagrawalAssesment4-28-03-25/internals/services"
	"omagrawalAssesment4-28-03-25/internals/utils"
	"sync"
)

var (
	wg = sync.WaitGroup{}
)

func main() {
	services.ReadDirFunc()
	services.CreatWorkerPool(&wg)
	go func() {
		wg.Wait()
		close(utils.FileDataOutput)
	}()
	services.ProcessAllData()
}
