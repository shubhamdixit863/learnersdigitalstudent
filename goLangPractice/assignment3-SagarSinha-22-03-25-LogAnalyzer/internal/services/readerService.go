package services

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func Reader(fileName string, chunkSize int, wg *sync.WaitGroup, results chan map[string]int) {
	var chunk []string

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		chunk = append(chunk, scanner.Text())
		if len(chunk) >= chunkSize {
			wg.Add(1)
			go Analyzer(chunk, results, wg)
			chunk = nil
		}
	}

	if len(chunk) > 0 {
		wg.Add(1)
		go Analyzer(chunk, results, wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

}
