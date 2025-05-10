package monitor

import (
	"filemonitor/internal/services"
	"filemonitor/internal/utils"
	"log"
	"sync"
)

type Monitor struct {
	fileChan chan string
	wg       *sync.WaitGroup
}

func NewMonitor(fileChan chan string, wg *sync.WaitGroup) *Monitor {
	return &Monitor{fileChan: fileChan, wg: wg}
}

func (m *Monitor) StartProcessing() {
	for range utils.ProcessLoopCount { 
		m.wg.Add(1)
		go func() {
			defer m.wg.Done()
			for filePath := range m.fileChan {
				err := services.ProcessFile(filePath)
				if err != nil {
					log.Printf(utils.FileProcessErr, filePath, err)
				}
			}
		}()
	}
}