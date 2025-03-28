package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

// FindTextFiles scans the directory and returns all .txt file paths
func FindTextFiles(dir string) ([]string, error) {
	var files []string

	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && filepath.Ext(path) == ".txt" {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}

// ReadFileLines reads each line of a file and sends it to a channel
func ReadFileLines(filename string, linesCh chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", filename, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linesCh <- scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", filename, err)
	}
}
