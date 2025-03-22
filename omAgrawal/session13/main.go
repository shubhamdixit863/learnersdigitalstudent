//package main
//
//import (
//	"fmt"
//	"log"
//	"sync"
//)
//
//func main() {
//	//func() {
//	//	for i := 0; i < 1000; i++ {
//	//		fmt.Println(i)
//	//		time.Sleep(1 * time.Second)
//	//	}
//	//
//	//}()
//	//time.Sleep(10 * time.Second)
//	//fmt.Println("hello")
//
//	var wg sync.WaitGroup
//
//	go func() {
//		log.Println("goroutine 1")
//		wg.Done()
//	}()
//	go func() {
//		fmt.Println("goroutine 2")
//		wg.Done()
//	}()
//	go func() {
//		fmt.Println("goroutine 3")
//		wg.Done()
//	}()
//	wg.Add(3)
//	wg.Wait()
//	log.Println("Now the wait is over and exited ")
//}

package main

import (
	"fmt"
	"sync"
)

func test1(num int, wg *sync.WaitGroup) {
	fmt.Println(2 * num)
	defer wg.Done()

}

func main() {
	var wg sync.WaitGroup
	numbers := []int{1, 2, 3, 4, 5}

	for _, v := range numbers {
		go test1(v, &wg)
	}
	wg.Add(5)
	wg.Wait()

	fmt.Println("exites sucesssfully")

}
