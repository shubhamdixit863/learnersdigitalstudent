package main

import (
	"fmt"
	"sync"
)

func multiply(num int, wg *sync.WaitGroup) {
	fmt.Printf("Processing %d --> %d \n", num, num*2)
	defer wg.Done()
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	var wg sync.WaitGroup
	wg.Add(5)
	for _, num := range numbers {
		go multiply(num, &wg)
	}

	wg.Wait()
	fmt.Println("All the goroutines completed")

}
