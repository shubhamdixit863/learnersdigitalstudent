package main

import (
	"log"
	"sync"
)

var c = 9

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 9; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			//Critical Section
			mu.Lock()
			c = c + 1
			mu.Unlock()
		}()
	}

	wg.Wait()

	log.Println(c)

}
