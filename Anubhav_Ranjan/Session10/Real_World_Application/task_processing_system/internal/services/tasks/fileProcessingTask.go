package tasks

import (
	"fmt"
	"task_processing_system/internal/services/errors"
	"task_processing_system/internal/utils"
)

type FileProcessingTask struct {
	FileName string
}

func NewFileProcessingTask(filename string) *FileProcessingTask {
	return &FileProcessingTask{FileName: filename}
}

func (fpt *FileProcessingTask) Run() error {
	content, err := utils.ReadFile(fpt.FileName)
	if err != nil {
		return errors.NewTaskError(fpt.Name(), err)
	}

	fmt.Println("Processing file:", fpt.FileName)
	fmt.Println("File Content:", content)
	return nil
}

func (fpt *FileProcessingTask) Name() string {
	return "FileProcessingTask"
}
