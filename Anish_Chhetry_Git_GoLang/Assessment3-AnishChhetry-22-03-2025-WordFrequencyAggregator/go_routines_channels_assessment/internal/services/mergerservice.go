package services

import (
	"sync"
)

type Merger struct {
	result map[string]int
	sync.Mutex
}

func NewMerger(res map[string]int) *Merger {
	return &Merger{result: res}
}

func (m *Merger) Merge(ch chan map[string]int) {

	for freqMap := range ch {
		m.Lock()
		for word, count := range freqMap {
			m.result[word] += count
		}
		m.Unlock()
	}
}
func (m *Merger) Result() map[string]int {
	return m.result
}
