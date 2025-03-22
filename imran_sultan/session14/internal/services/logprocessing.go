package services

import (
	"fmt"
	"sync"
	"time"
)

func LogProcessing(ch chan string, newch chan *LogEntry, wg *sync.WaitGroup) {
	defer wg.Done()

	val := <-ch
	T := val[:27]
	Level := val[52:56]
	Message := val[58:]

	layout := "2006-01-02 15:04:05.999999999"
	t, err := time.Parse(layout, T)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}
	obj := NewLog(t, Level, Message)
	newch <- obj
	close(newch)

}
