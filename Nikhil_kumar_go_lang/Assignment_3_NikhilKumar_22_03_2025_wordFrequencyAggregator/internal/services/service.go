package services

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"wordFrequencyAggregator/internal/models"
	"wordFrequencyAggregator/internal/utils"
)

type WordFrequency map[string]int

func ReadFileInChunks(filePath string, chunkSize int) ([]models.Chunk, error) {
	log.Println("Reading file...")
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf(utils.FAILD_TO_OPEN, err)
	}
	defer file.Close()

	var chunks []models.Chunk
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var currentChunk strings.Builder
	chunkID := 0
	lineCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		currentChunk.WriteString(line + "\n")
		lineCount++

		if lineCount >= chunkSize {
			chunks = append(chunks, models.Chunk{
				Data:    currentChunk.String(),
				ChunkID: chunkID,
			})
			currentChunk.Reset()
			lineCount = 0
			chunkID++
		}
	}

	if currentChunk.Len() > 0 {
		chunks = append(chunks, models.Chunk{
			Data:    currentChunk.String(),
			ChunkID: chunkID,
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf(utils.READ_ERR, err)
	}

	return chunks, nil
}

func MergeWordFrequency(frequencies []WordFrequency, mutex *sync.Mutex) WordFrequency {
	merged := make(WordFrequency)

	for _, freq := range frequencies {
		mutex.Lock()
		for word, count := range freq {
			merged[word] += count
		}
		mutex.Unlock()
	}

	return merged
}
