package main

import (
	"log"
	"sync"
)

func gr(number int, wg *sync.WaitGroup) {
	log.Printf("Processing %d------> %d\n", number, number*2)
	wg.Done()

}


func main() {
	var wg sync.WaitGroup
	numbers := []int{1, 2, 3, 4, 5}
	wg.Add(len(numbers))
	for _, n := range numbers {
		go gr(n, &wg)
	}
	wg.Wait()
	log.Println("All grs completed!")

}