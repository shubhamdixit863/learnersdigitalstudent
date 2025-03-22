package services

import (
	"fmt"
	"sync"
)

func Store(send chan *LogEntry, wq *sync.WaitGroup) {
	defer wq.Done()
	for receive := range send {
		Loglist = append(Loglist, receive)
		Mapp[receive.Level] = Mapp[receive.Level] + 1
		fmt.Println(receive.Level, Mapp[receive.Level])
	}
}
