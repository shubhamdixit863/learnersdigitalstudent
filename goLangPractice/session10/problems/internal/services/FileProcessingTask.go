package services

import (
	"fmt"
	"os"
	"session10/problems/internal/utils"
)

type FileProcessingTask struct {
	Filename string
}

func NewFileProcessingTask(filename string) *FileProcessingTask {
	return &FileProcessingTask{
		Filename: filename,
	}
}

func (f FileProcessingTask) Run() error {
	defer fmt.Println("FileProcessingTask: Cleanup")

	fmt.Println("Reading file data...")

	data, err := os.ReadFile(f.Filename)
	if err != nil {
		return fmt.Errorf("%w: Unable to open file", utils.TaskError)
	}

	fmt.Printf("File contents: %s\n", data)
	return nil
}
