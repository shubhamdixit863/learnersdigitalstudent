package main

import (
	"fmt"
	"session14/sorting"
	"sync"
)

// Shared Implementation
//var c = 9
//
//func main() {
//	var wg sync.WaitGroup
//
//	go func() {
//		defer wg.Done()
//		c = 8
//	}()
//
//	go func() {
//		defer wg.Done()
//		fmt.Println("Exiting...", c)
//		if c > 9 {
//			fmt.Println()
//		}
//		fmt.Println("Exiting...")
//	}()
//
//	wg.Add(2)
//
//	wg.Wait()
//
//}
//	 This is a worng way of implementation

func main() {
	//var wg sync.WaitGroup
	//
	//ch := make(chan int)
	//
	//go func() {
	//	defer wg.Done()
	//	ch <- 9
	//	close(ch) // if not closed, it is the waste of resource
	//}()
	//
	//go func() {
	//	defer wg.Done()
	//	y := <-ch
	//	fmt.Println("Data from channel: ", y)
	//}()
	//
	//wg.Add(2)
	//wg.Wait()

	//var slc = []int{8, 1, 3, 90, 2}
	//fmt.Println("Before Sorting: ", slc)
	//sorting.BubbleSort(slc)
	//fmt.Println("After Sorting: ", slc)

	var wg sync.WaitGroup
	var slc = []int{8, 1, 3, 90, 2}
	//ch := make(chan int)
	chbuf := make(chan int, 5)

	go sorting.BubbleOuterBuf(slc, chbuf, &wg)
	go sorting.BubbleInnerBuf(slc, chbuf, &wg)

	wg.Add(2)
	wg.Wait()

	fmt.Println("After Sorting: ", slc)
}
