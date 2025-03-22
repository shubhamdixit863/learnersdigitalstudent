package services

import (
	"fmt"
	"math/rand"
	"session16/storage"
	"sync"
)

var m sync.Mutex

func ProcessingOrder(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	item := <-ch
	if storage.MpItem[item] > 0 {
		m.Lock()
		storage.MpItem[item]--
		m.Unlock()
		fmt.Println("Shipped")
		fmt.Println("your order")
		fmt.Println("#", rand.Int())
		fmt.Println("Completed")
	} else {
		fmt.Println("item no available!")
	}

}
