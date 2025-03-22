package services

import (
	"fmt"
	"math/rand"
	"session16/storage"
	"sync"
)

func Order(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	indx := rand.Intn(5)
	item := storage.ListItem[indx]
	fmt.Println("Processing...")
	ch <- item

}
