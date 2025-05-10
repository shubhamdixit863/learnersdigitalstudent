package wordprocessor

import (
	"bufio"
	"fmt"
	"strings"
	"sync"
	"wordFrequencyAggregator/internal/models"
	"wordFrequencyAggregator/internal/services"
	"wordFrequencyAggregator/internal/utils"
)

type Processor interface {
	Process(chunk models.Chunk) (services.WordFrequency, error)
}

type WordProcessor struct{}

func (wp WordProcessor) Process(chunk models.Chunk) (services.WordFrequency, error) {
	freq := make(services.WordFrequency)
	scanner := bufio.NewScanner(strings.NewReader(chunk.Data))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		freq[word]++
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf(utils.SCANNER_ERR, chunk.ChunkID, err)
	}

	return freq, nil
}

func ChunkProcessor(chunk models.Chunk, processor Processor, results chan<- services.WordFrequency, wg *sync.WaitGroup, errChan chan<- error) {
	defer wg.Done()

	freq, err := processor.Process(chunk)
	if err != nil {
		errChan <- err
		return
	}

	results <- freq
}
