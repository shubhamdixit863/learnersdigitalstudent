package main

import (
	"fmt"
	"log"
	"parallel_file_hasher/filewalker"
	"parallel_file_hasher/hasher"
	"parallel_file_hasher/results"
)

func main() {
	files, err := filewalker.WalkDirectory("./files")
	if err != nil {
		log.Fatal(err)
	}

	resultChan := make(chan results.HashResult, len(files))

	for _, file := range files {
		go hasher.ComputeHash(file, resultChan)
	}

	for range files {
		hashResult := <-resultChan
		fmt.Printf("File: %s | Hash: %s\n", hashResult.File, hashResult.Hash)
	}

	close(resultChan)
}
