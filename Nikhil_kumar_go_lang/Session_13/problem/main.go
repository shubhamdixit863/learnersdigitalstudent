package main

import (
	"log"
	"sync"
)

func Multi_by_two(a int, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Println("Processing ", a, "->", a*2)
}

func main() {
	var wg sync.WaitGroup
	number := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(number); i++ {
		go Multi_by_two(number[i], &wg)
	}

	wg.Add(5)
	wg.Wait()

	log.Println("All goroutines completed!")
}
