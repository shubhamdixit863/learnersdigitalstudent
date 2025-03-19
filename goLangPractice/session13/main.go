package main

import (
	"log"
	"sync"
)

func foo() {
	log.Println("Hey there")
}

//func main() {
//	//go func() {
//	//	for i := 0; i < 1000; i++ {
//	//		fmt.Println(i)
//	//		time.Sleep(1 * time.Second)
//	//	}
//	//	fmt.Println("Hello")
//	//}()
//	//
//	//go foo()
//	//
//	//time.Sleep(10 * time.Second) // Making main go routine sleep for 10 second
//	////fmt.Println("Main Routine exits")
//	//fmt.Scanln() // This also works but should not be used
//
//	// Both of this time and scan method are bad ways better way.So, we use syncWait
//
//}

func main() {
	var wg sync.WaitGroup

	go func() {
		log.Println("GoRoutine 1")
		wg.Done() // it is decrementing the wg.Add counter
	}()

	go func() {
		log.Println("GoRoutine 2")
		wg.Done()
	}()

	go func() {
		log.Println("GoRoutine 3")
		wg.Done()
	}()

	wg.Add(3) // Make a counter of no. of go routines

	wg.Wait() // will actually holds the go routine from exiting

	log.Println("Now the wait is over exiting the go routine")
}
