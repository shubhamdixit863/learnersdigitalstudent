package main

import (
	"log"
	"sync"
)

func goroutine1(number int, wg *sync.WaitGroup) {
	//log.Println("goroutine 1")
	log.Printf("Processing %d------> %d\n", number, number*2)
	wg.Done()

}
func goroutine2(number int, wg *sync.WaitGroup) {
	//log.Println("goroutine 2")
	log.Printf("Processing %d------> %d\n", number, number*2)
	wg.Done()

}

func main() {
	var wg sync.WaitGroup
	numbers := []int{1, 2, 3, 4, 5}
	for _, n := range numbers {
		go goroutine1(n, &wg)
		go goroutine2(n, &wg)
	}
	wg.Add(len(numbers))
	wg.Wait()

	log.Println("All goroutines completed!")

}
