package main

import (
	"fmt"
	"go_routines_channels_assessment/internal/services"
	"go_routines_channels_assessment/internal/utils"
	"os"
	"strings"
	"sync"
)

const (
	fileName    = "book.txt"
	printResult = "%s: %d\n"
)

func main() {
	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(utils.FileOpeningError(fileName))
		return
	}

	text := string(content)
	lines := strings.Split(text, "\n")

	ch := make(chan map[string]int)
	wordCounter := services.NewSimpleWordCounter()
	chunkProcessor := services.NewChunkProcessor(wordCounter)

	var wg sync.WaitGroup

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		wg.Add(1)
		go chunkProcessor.ProcessChunk(line, &wg, ch)
	}

	merger := services.NewMerger(make(map[string]int))

	go func() {
		wg.Wait()
		close(ch)
	}()

	merger.Merge(ch)

	PrintResult(merger)
}

func PrintResult(merger *services.Merger) {
	for word, count := range merger.Result() {
		fmt.Printf(printResult, word, count)
	}
}
