package services

import (
	"bufio"
	"concurrentFileProcessor/internal/utils"
	"log"
	"os"
	"path/filepath"
	"sync"
)

func DirectoryToFiles(directoryPath string) []string {
	files, err := filepath.Glob(filepath.Join(directoryPath, utils.TXT))
	if err != nil {
		log.Fatal(err)
		return nil
	}
	if len(files) == 0 {
		log.Println(utils.NO_TXT_FOUND)
		return nil
	}
	return files
}

func ReadLines(filePath string, linesChan chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open(filePath)
	if err != nil {
		log.Printf(utils.ERR_OPENING_FILE, filePath, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linesChan <- scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Printf(utils.ERR_READING_FILE, filePath, err)
	}
}
