package processor

import (
	"assessment3/internal/types"
	"strconv"
)

type Processor interface {
	Process(chunk [][2]string) types.UserSpend
}

type ChunkProcessor struct{}

func NewChunkProcessor() *ChunkProcessor {
	return &ChunkProcessor{}
}

func (cp *ChunkProcessor) Process(chunk [][2]string) types.UserSpend {
	userSpend := make(types.UserSpend)
	for _, record := range chunk {
		userID := record[0]
		amount, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			continue // Skip invalid entries
		}
		userSpend[userID] += amount
	}
	return userSpend
}
