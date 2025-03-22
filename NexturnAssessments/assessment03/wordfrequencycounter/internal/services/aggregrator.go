package services

import "sync"

var mu sync.Mutex
func AggregateResults(resultCh <-chan map[string]int, finalFreq map[string]int, doneCh chan<- bool) {
	for freqMap := range resultCh {
		mu.Lock()
		for word, count := range freqMap {
			finalFreq[word] += count
		}
		mu.Unlock()
	}
	doneCh <- true
}
