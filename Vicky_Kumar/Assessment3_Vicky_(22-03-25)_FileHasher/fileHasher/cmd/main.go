package main

import (
	"fmt"
	"log"
	"sync"

	"fileHasher/internal/filewalker"
	"fileHasher/internal/hasher"
	"fileHasher/internal/model"
	"fileHasher/internal/utils"
)

const dirPath = "./test_files"

func main() {
	// files := []string{
	// 	"test_files/file1.txt",
	// 	"test_files/file2.txt",
	// 	"test_files/file3.txt",
	// }

	fileWalker := filewalker.NewFileWalker()
	files, err := fileWalker.ReadDirectory(dirPath)
	if err != nil {
		log.Println(err)
	}

	if len(files) == 0 {
		log.Println(utils.NewCustomError(utils.ErrNoFilesFound))
	}

	resultChan := make(chan model.HashResult, len(files))
	var wg sync.WaitGroup

	fileHasher := hasher.NewFileHasher()

	for _, filePath := range files {
		wg.Add(1)
		go func(path string) {
			defer wg.Done()

			hash, err := fileHasher.HashFile(path)
			if err != nil {
				resultChan <- *model.NewHashRsult(path, err.Error())
				return
			}
			resultChan <- *model.NewHashRsult(path, hash)
		}(filePath)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	var results []model.HashResult
	for result := range resultChan {
		results = append(results, result)
	}

	displayResults(results)
}

func displayResults(results []model.HashResult) {
	fmt.Println("File Hashes:")
	fmt.Println()
	for _, result := range results {
		fmt.Printf("%s: %s\n", result.File, result.Hash)
	}
}
