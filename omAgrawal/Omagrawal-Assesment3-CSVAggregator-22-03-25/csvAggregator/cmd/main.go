package main

import (
	"csvAggregator/internals/models"
	"csvAggregator/internals/services"
	"fmt"
	"log"
	"sync"
)

var (
	Ch        = make(chan map[string]float64, models.N)
	FinalMapp = make(map[string]float64)
	wg        = sync.WaitGroup{}
	mu        = sync.Mutex{}
)

func main() {

	Filedata := services.FileReader()[1:]

	for key, value := range Filedata {
		fmt.Println(key, "  ", value)
	}

	services.SplitFileToN(Filedata)

	for _, v := range models.SplitFileMapp {
		wg.Add(1)
		go services.ProcessData(Ch, v, &wg)
	}

	wg.Wait()
	close(Ch)
	for i := 0; i < models.N; i++ {
		mu.Lock()
		mpp := <-Ch

		for i, v := range mpp {
			FinalMapp[i] = FinalMapp[i] + v
		}
		mu.Unlock()
	}

	log.Println(models.FinalOutput)
	for key, value := range FinalMapp {
		fmt.Println(key, "  ", value)
	}

}
