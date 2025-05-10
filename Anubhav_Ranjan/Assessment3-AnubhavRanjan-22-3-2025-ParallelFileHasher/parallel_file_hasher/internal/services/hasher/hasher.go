package hasher

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"sync"

	"parallel_file_hasher/internal/models"
)

const (
	errOpenFileFmt = "failed to open file: %w"
	errReadFileFmt = "failed to read file: %w"
	errComputeHash = "failed to compute hash: %w"
)

type HashStrategy interface {
	Hash(data []byte) (string, error)
}

type SHA256Strategy struct{}

func (s *SHA256Strategy) Hash(data []byte) (string, error) {
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:]), nil
}

type HasherService struct {
	workers      int
	hashStrategy HashStrategy
}

func NewHasherService(workers int) *HasherService {
	return &HasherService{
		workers:      workers,
		hashStrategy: &SHA256Strategy{},
	}
}

func (h *HasherService) SetHashStrategy(strategy HashStrategy) {
	h.hashStrategy = strategy
}

func (h *HasherService) ProcessFiles(files []string) ([]models.HashResult, error) {
	var wg sync.WaitGroup
	resultChan := make(chan models.HashResult, len(files))

	fileChan := make(chan string, h.workers)

	for i := 0; i < h.workers; i++ {
		wg.Add(1)
		go h.worker(fileChan, resultChan, &wg)
	}

	for _, file := range files {
		fileChan <- file
	}
	close(fileChan)

	wg.Wait()
	close(resultChan)

	var results []models.HashResult
	for result := range resultChan {
		results = append(results, result)
	}

	return results, nil
}

func (h *HasherService) worker(fileChan <-chan string, resultChan chan<- models.HashResult, wg *sync.WaitGroup) {
	defer wg.Done()

	for file := range fileChan {
		result := h.hashFile(file)
		resultChan <- result
	}
}

func (h *HasherService) hashFile(file string) models.HashResult {
	result := models.HashResult{
		File: file,
	}

	f, err := os.Open(file)
	if err != nil {
		result.Error = fmt.Errorf(errOpenFileFmt, err)
		return result
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		result.Error = fmt.Errorf(errReadFileFmt, err)
		return result
	}

	hash, err := h.hashStrategy.Hash(data)
	if err != nil {
		result.Error = fmt.Errorf(errComputeHash, err)
		return result
	}

	result.Hash = hash
	return result
}
