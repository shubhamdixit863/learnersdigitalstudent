package main

import (
	"fmt"
	"sync"
)

//	func main() {
//		go func() {
//			fmt.Println("hey baby")
//		}()
//		time.Sleep(10 * time.Second) // bad way of doing things
//		fmt.Println("exit")
//	}
func goroutine1(num int, wg *sync.WaitGroup) {
	fmt.Println("Processing ", num, "-->", num*2)
	wg.Done()
}
func goroutine2(num int, wg *sync.WaitGroup) {
	fmt.Println("Processing ", num, "-->", num*2)
	wg.Done()
}
func goroutine3(num int, wg *sync.WaitGroup) {
	fmt.Println("Processing ", num, "-->", num*2)
	wg.Done()
}
func goroutine4(num int, wg *sync.WaitGroup) {
	fmt.Println("Processing ", num, "-->", num*2)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	numbers := []int{1, 2, 3, 4}

	go goroutine1(numbers[0], &wg)
	go goroutine2(numbers[1], &wg)
	go goroutine3(numbers[2], &wg)
	go goroutine4(numbers[3], &wg)

	wg.Add(4)
	wg.Wait()

	fmt.Println("exit..!")
}
