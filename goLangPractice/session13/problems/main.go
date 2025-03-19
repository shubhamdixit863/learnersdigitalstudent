package main

import (
	"fmt"
	"sync"
)

func multiply(a *int, wg *sync.WaitGroup) {
	result := *a * 2
	fmt.Println("Processing ", *a, " -> ", result)

	wg.Done()
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	var wg sync.WaitGroup

	wg.Add(len(numbers))

	for _, number := range numbers {
		go multiply(&number, &wg)
	}

	wg.Wait()

	fmt.Println("All goroutines completed!")
}
