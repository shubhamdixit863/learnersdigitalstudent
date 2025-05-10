package services

import (
	"errors"
	"fmt"
	"log"
	"os"
)

type FileProcessingTask struct {
	file string
}

func (s *FileProcessingTask) Run() error {

	file, err := os.Open(s.file)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("%s file not found", s.file)
		}
	}
	log.Println("File opened: ", s.file)

	defer file.Close()
	return nil
}
func (s *FileProcessingTask) ReturnName() string {
	return "FileProcessingTask"
}

func NewFileProcessingTask(file string) *FileProcessingTask {

	return &FileProcessingTask{file: file}
}
