package services

import (
	"fmt"
	"sync"
)

func LogProcessing(log chan string, send chan *LogEntry, wg *sync.WaitGroup) {
	defer wg.Done()

	for logval := range log {
		fmt.Println(logval)
		tim := logval[:27]
		er := logval[53:58]
		loggg := logval[59:]
		//fmt.Println(tim)
		//fmt.Println(er)
		//fmt.Println(loggg)

		v := NewLog(tim, er, loggg)
		fmt.Println(v)
		send <- v
	}
	close(send)

}
