package main

import (
	"fmt"
	"sync"
)

func Multiply(wg *sync.WaitGroup, num int) {
	fmt.Printf("Processing %d -> %d \n", num, 2*num)
	wg.Done()
}

func main() {
	wg := new(sync.WaitGroup)
	numbers := []int{1, 2, 3, 4, 5}

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go Multiply(wg, numbers[i])
	}
	wg.Wait()
	fmt.Println("All goroutines completed!")

}
