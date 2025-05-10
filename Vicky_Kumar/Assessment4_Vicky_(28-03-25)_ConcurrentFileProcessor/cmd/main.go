package main

import (
	"fmt"
	"log"
	"sync"

	"concurrent_file_processor/internal/filewalker"
	"concurrent_file_processor/internal/services"
	"concurrent_file_processor/internal/utils"
)

const dirPath = "./test_files"

var processMode string
var keyword string = utils.EmptyString

func main() {

	fileWalker := filewalker.NewFileWalker()
	// reading all files in a directory
	files, err := fileWalker.ReadDirectory(dirPath)
	if err != nil {
		log.Println(err)
	}

	if len(files) == 0 {
		log.Println(utils.NewCustomError(utils.ErrNoFilesFound))
		return
	}
	// readling lines of all files and sending to linesChan
	linesChan := make(chan string, 100)
	resultChan := make(chan string, 100)
	var wg sync.WaitGroup

	// fan out
	for _, file := range files {
		wg.Add(1)
		go func(filePath string) {
			defer wg.Done()
			fileWalker.ReadFileLines(filePath, linesChan)
		}(file)
	}

	go func() {
		wg.Wait()
		close(linesChan)
	}()

	handleInput()

	p := services.NewProcessor(processMode, keyword)

	wg.Add(1)
	go p.Process(linesChan, resultChan, &wg)
	wg.Wait()
	for result := range resultChan {
		log.Println(result)
	}

}

func handleInput() {
	for _, option := range utils.PrintOptions {
		log.Println(option)
	}

	var c int
	fmt.Scanln(&c)

	switch c {
	case 1:
		processMode = utils.Options[0]
		log.Println(utils.KeyWord)
		fmt.Scanln(&keyword)
	case 2:
		processMode = utils.Options[1]
	case 3:
		processMode = utils.Options[2]
	case 4:
		processMode = utils.Options[3]
	default:
		log.Println(utils.WordCount)
	}
}
