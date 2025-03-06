package services

import (
	"errors"
	"fmt"
	"log"
	"os"
	"practical2/internal/utils"
)

type DataValidationTask struct {
	file string
}

func (s *DataValidationTask) Run() error {
	file, err := os.ReadFile(s.file)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("cant open file %s", s.file)
		}
	}
	if len(file) == 0 {
		return fmt.Errorf("empty file: %s \n %v", s.file, utils.NewValidationError())
	}
	log.Println("File Read, content:\n", string(file))
	return nil
}
func (s *DataValidationTask) ReturnName() string {
	return "DataValidationTask"
}
func NewDataValidationTask(file string) *DataValidationTask {
	return &DataValidationTask{file: file}
}
