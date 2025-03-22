package main

import (
	"session16/services"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	ch := make(chan string)

	// indx := rand.Intn(5)
	// item := storage.ListItem[indx]
	for i := 0; i < 5; i++ {
		go services.Order(ch, &wg)
		go services.ProcessingOrder(ch, &wg)

		wg.Add(2)
	}

	wg.Wait()

	close(ch)
	// fmt.Println(storage.MpItem)
}
