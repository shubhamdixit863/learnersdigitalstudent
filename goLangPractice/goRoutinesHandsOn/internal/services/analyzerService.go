package services

import (
	"strings"
	"sync"
)

func Analyzer(lines []string, results chan map[string]int, wg *sync.WaitGroup) {
	defer wg.Done()

	logCounts := make(map[string]int)

	for _, line := range lines {
		if strings.Contains(line, "INFO") {
			logCounts["INFO"]++
		} else if strings.Contains(line, "ERROR") {
			logCounts["ERROR"]++
		} else if strings.Contains(line, "DEBUG") {
			logCounts["DEBUG"]++
		} else if strings.Contains(line, "WARNING") {
			logCounts["WARNING"]++
		}

	}

	results <- logCounts
}
