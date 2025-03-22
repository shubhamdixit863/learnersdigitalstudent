package services

import (
	"bufio"
	"fmt"
	"os"
)

const fileopenerror="failed to open file: %w"
const filereaderror="failed to read file:%v"

func ReadFile(filePath string, chunkSize int, ch chan<- []byte) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf(fileopenerror, err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		buffer := make([]byte, chunkSize)
		n, err := reader.Read(buffer)
		if n > 0 {
			ch <- buffer[:n]
		}
		if err != nil {
			if err.Error() != "EOF" {
				return fmt.Errorf(filereaderror, err)
			}
			break
		}
	}
	close(ch)
	return nil
}
