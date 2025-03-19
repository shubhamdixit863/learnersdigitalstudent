package status

import (
	"fmt"
	"sync"
)

var Queue = make(chan string, 100)

func StatusLogger(wg *sync.WaitGroup) {
	defer wg.Done()
	for status := range Queue {
		fmt.Println(status)
	}
}
