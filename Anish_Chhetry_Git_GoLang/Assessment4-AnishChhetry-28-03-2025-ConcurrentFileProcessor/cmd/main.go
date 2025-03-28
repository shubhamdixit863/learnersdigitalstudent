package main

import (
	"golang_assessment_4/internal/services"
	"log"
	"os"
	"sync"
)

const (
	dirLoc    = "../textfiles"
	weightAdd = 1
)

func main() {
	files, err := os.ReadDir(dirLoc)
	if err != nil {
		log.Println(err)
	}

	fileNamesChan := make(chan string, len(files))
	contentChan := make(chan map[string]string, len(files))

	var wg sync.WaitGroup
	var storage = make(map[string]string)

	for _, file := range files {
		wg.Add(weightAdd)
		fileNamesChan <- dirLoc + "/" + file.Name()
		go func() {
			defer wg.Done()
			filePath := <-fileNamesChan
			content := services.ReadFile(filePath)
			contentChan <- map[string]string{filePath: content}
		}()
	}

	close(fileNamesChan)
	go func() {
		wg.Wait()
		close(contentChan)
	}()

	for content := range contentChan {
		for k, v := range content {
			storage[k] = v
		}
	}
	services.User(storage)
}
