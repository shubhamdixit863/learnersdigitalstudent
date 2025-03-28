package fileprocessor

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"sync"

	"concurrent_file_processor/internal/config"
	"concurrent_file_processor/internal/errors"
	"concurrent_file_processor/internal/utils"
)

type FileProcessor struct {
	config   *config.Config
	strategy ProcessingStrategy
	mu       sync.Mutex
	fileMap  map[string]bool
}

type ProcessResult struct {
	Filename string
	Results  []string
	Error    error
}

func NewFileProcessor(cfg *config.Config) (*FileProcessor, error) {
	strategy, err := StrategyFactory(cfg)
	if err != nil {
		return nil, err
	}

	return &FileProcessor{
		config:   cfg,
		strategy: strategy,
		fileMap:  make(map[string]bool),
	}, nil
}

func (fp *FileProcessor) Process(ctx context.Context) ([]ProcessResult, error) {
	// Find text files
	files, err := utils.FindTextFiles(fp.config.Directory)
	if err != nil {
		return nil, fmt.Errorf(errors.ErrInvalidDirectoryMsg, err)
	}

	if len(files) == 0 {
		return nil, fmt.Errorf(errors.ErrNoTextFilesMsg)
	}

	rawLinesChan := make(chan string, 10)
	processedLinesChan := make(chan string, 10)
	resultsChan := make(chan ProcessResult, len(files))

	var producerWg sync.WaitGroup
	var processorWg sync.WaitGroup

	distributedFiles := fp.distributeFiles(files)

	for _, fileGroup := range distributedFiles {
		producerWg.Add(1)
		go fp.fileProducer(ctx, fileGroup, rawLinesChan, &producerWg)
	}

	go func() {
		producerWg.Wait()
		close(rawLinesChan)
	}()

	numWorkers := fp.config.MaxGoroutines
	for i := 0; i < numWorkers; i++ {
		processorWg.Add(1)
		go fp.lineProcessor(ctx, rawLinesChan, processedLinesChan, &processorWg)
	}

	go func() {
		processorWg.Wait()
		close(processedLinesChan)
	}()

	go fp.resultAggregator(ctx, processedLinesChan, resultsChan)

	var finalResults []ProcessResult
	for result := range resultsChan {
		finalResults = append(finalResults, result)
	}

	return finalResults, nil
}

func (fp *FileProcessor) distributeFiles(files []string) [][]string {
	fp.mu.Lock()
	defer fp.mu.Unlock()

	fp.fileMap = make(map[string]bool)

	numProducers := min(len(files), fp.config.MaxGoroutines)
	distributedFiles := make([][]string, numProducers)

	for i, file := range files {
		if _, processed := fp.fileMap[file]; !processed {
			producerIndex := i % numProducers
			distributedFiles[producerIndex] = append(distributedFiles[producerIndex], file)
			fp.fileMap[file] = true
		}
	}

	return distributedFiles
}

func (fp *FileProcessor) fileProducer(
	ctx context.Context,
	files []string,
	rawLinesChan chan<- string,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	for _, filename := range files {
		select {
		case <-ctx.Done():
			return
		default:
			if err := fp.readFile(ctx, filename, rawLinesChan); err != nil {
				fmt.Printf(errors.ErrReadingFileMsg, filename, err)
			}
		}
	}
}

func (fp *FileProcessor) readFile(
	ctx context.Context,
	filename string,
	rawLinesChan chan<- string,
) error {
	fp.mu.Lock()
	file, err := os.Open(filename)
	fp.mu.Unlock()

	if err != nil {
		return fmt.Errorf(errors.ErrProcessingFailedMsg,
			fmt.Sprintf(errors.ErrOpenFile, filename, err))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return nil
		case rawLinesChan <- scanner.Text():
		}
	}

	return scanner.Err()
}

func (fp *FileProcessor) lineProcessor(
	ctx context.Context,
	rawLinesChan <-chan string,
	processedLinesChan chan<- string,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	for line := range rawLinesChan {
		select {
		case <-ctx.Done():
			return
		default:
			result, err := fp.strategy.Process(line)
			if err != nil {
				fmt.Printf(errors.ErrProcessingFmt, err)
				continue
			}
			if result != "" {
				processedLinesChan <- result
			}
		}
	}
}

func (fp *FileProcessor) resultAggregator(
	ctx context.Context,
	processedLinesChan <-chan string,
	resultsChan chan<- ProcessResult,
) {
	defer close(resultsChan)

	results := make(map[string][]string)

	for line := range processedLinesChan {
		select {
		case <-ctx.Done():
			return
		default:
			results["processed"] = append(results["processed"], line)
		}
	}

	resultsChan <- ProcessResult{
		Filename: "Aggregated Results",
		Results:  results["processed"],
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
