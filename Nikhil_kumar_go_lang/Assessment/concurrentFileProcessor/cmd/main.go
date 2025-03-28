package main

import (
	"concurrentFileProcessor/internal/services"
	"concurrentFileProcessor/internal/utils"
	"concurrentFileProcessor/mode"
	"fmt"
	"log"
	"sync"
)

func main() {
	log.Println(utils.TITLE)
	log.Print(utils.INPUT_DIR)

	var directoryPath string
	fmt.Scanln(&directoryPath)

	files := services.DirectoryToFiles(directoryPath)

	var wg sync.WaitGroup
	linesChan := make(chan string, 100)

	for _, file := range files {
		wg.Add(1)
		go services.ReadLines(file, linesChan, &wg)
		log.Println(utils.PROC_FILE, file)
	}

	go func() {
		wg.Wait()
		close(linesChan)
	}()

	mode.Mode(linesChan)
}
