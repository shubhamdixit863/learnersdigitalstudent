package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock() // Locking the counter
			counter++
			mu.Unlock() // Unlocking after update
		}()
	}
	fmt.Println("Counter:", counter)
	wg.Wait()
	fmt.Println("Final Counter:", counter)

	defer wait()
	usageDefer()

}

//

func wait() {
	var wg sync.WaitGroup
	// Creating multiple goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d working…\n", id)
		}(i)
	}
	wg.Wait()
	fmt.Println("All goroutines finished!")
}

func usageDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)

	}
	fmt.Println("The numbers are:")
}
