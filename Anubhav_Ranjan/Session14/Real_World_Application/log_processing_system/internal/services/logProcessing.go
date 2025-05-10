package services

import (
	"fmt"
	"regexp"
	"sync"
	"time"
)

type LogProcessor struct {
	id         int
	inputChan  <-chan string
	outputChan chan<- LogEntry
	wg         *sync.WaitGroup
	logPattern *regexp.Regexp
}

func NewLogProcessor(id int, inputChan <-chan string, outputChan chan<- LogEntry, wg *sync.WaitGroup) *LogProcessor {
	pattern := regexp.MustCompile(`\[([\d-]+ [\d:]+)\] ([A-Z]+): (.+)`)

	return &LogProcessor{
		id:         id,
		inputChan:  inputChan,
		outputChan: outputChan,
		wg:         wg,
		logPattern: pattern,
	}
}

func (lp *LogProcessor) Process() {
	defer lp.wg.Done()

	for logString := range lp.inputChan {
		matches := lp.logPattern.FindStringSubmatch(logString)

		fmt.Println("Matches:", matches[0])

		if len(matches) == 4 {
			timestampStr := matches[1]
			level := matches[2]
			message := matches[3]

			t, err := time.Parse("2006-01-02 15:04:05", timestampStr)
			if err != nil {
				fmt.Printf("Processor %d: Error parsing timestamp: %v\n", lp.id, err)
				continue
			}

			entry := LogEntry{
				Timestamp: t,
				Level:     level,
				Message:   message,
			}

			lp.outputChan <- entry
		} else {
			fmt.Printf("Processor %d: Failed to parse log: %s\n", lp.id, logString)
		}
	}

	fmt.Printf("Processor %d completed\n", lp.id)
}
