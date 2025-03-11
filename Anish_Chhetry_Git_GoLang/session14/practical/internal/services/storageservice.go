package services

import (
	"sync"
)

var info, err, warn = 0, 0, 0

var FinalMap = make(map[string]int)

func StorageService(ch chan LogEntry, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		if v.Level == "INFO" {
			info++
		} else if v.Level == "WARN" {
			warn++
		} else if v.Level == "ERROR" {
			err++
		}
	}
	FinalMap["INFO"] = info
	FinalMap["ERROR"] = err
	FinalMap["WARN"] = warn

	//fmt.Println("Log Level Count Summary: INFO", info, "WARNING", warn, "ERROR", err)

}
