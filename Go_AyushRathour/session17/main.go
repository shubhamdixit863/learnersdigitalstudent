package main

import (
	"fmt"
	"strings"
	"sync"
)

func worker(taskCh chan string, wg *sync.WaitGroup, workerId int, resultCh chan string) {
	// Any job
	defer wg.Done()
	for v := range taskCh {
		s := fmt.Sprintf("Workerid %d - %s\n", workerId, strings.ToUpper(v))
		resultCh <- s
	}

}

const numWorkers = 2

func main() {
	var wg sync.WaitGroup
	taskChan := make(chan string)
	resultChan := make(chan string)
	fn := []string{"a.txt", "b.txt", "c.txt", "d.txt"}

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(taskChan, &wg, i, resultChan)
	}

	//wg.Add(1)
	go func() {
		for g := range resultChan {
			fmt.Println(g)
		}
		//defer wg.Done()
	}()

	for i := 0; i < len(fn); i++ {
		taskChan <- fn[i]
	}

	close(taskChan)
	wg.Wait()
	close(resultChan)

}
