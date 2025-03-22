package services

import (
	"go_routines_channels_assessment/internal/model"
	"sync"
)

type ChunkProcessor struct {
	counter model.WordCounter
}

func NewChunkProcessor(wordCounter model.WordCounter) *ChunkProcessor {
	return &ChunkProcessor{
		counter: wordCounter,
	}
}

func (cp ChunkProcessor) ProcessChunk(chunk string, wg *sync.WaitGroup, ch chan map[string]int) {
	defer wg.Done()
	finalMap := cp.counter.CountWords(chunk)
	ch <- finalMap
}
