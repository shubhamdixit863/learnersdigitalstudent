package services

import (
	"goConcurrencyHandsOn/internal/utilities"
	"log"
	"os"
	"sync"
)

func ScanDirectory(path string, lines chan string, wg *sync.WaitGroup) {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Println(utilities.ErrorDirOpen, err)
	}

	for _, file := range files {
		filePath := path + utilities.Slash + file.Name()
		wg.Add(1)
		go ReadLines(filePath, lines, wg)
	}

	go func() {
		wg.Wait()
		close(lines)
	}()
}
