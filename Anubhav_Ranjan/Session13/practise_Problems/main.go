package main

import (
	"fmt"
	"sync"
)

func processNumber(num int, wg *sync.WaitGroup) {
	defer wg.Done()

	result := num * 2

	fmt.Printf("Processing %d -> %d\n", num, result)
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	var wg sync.WaitGroup

	wg.Add(len(numbers))

	for _, num := range numbers {
		go processNumber(num, &wg)
	}

	wg.Wait()

	fmt.Println("All goroutines completed!")
}
