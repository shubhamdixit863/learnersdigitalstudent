package main

import (
	"concurrent_file_processing/internals/services"
	"concurrent_file_processing/internals/utils"
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
