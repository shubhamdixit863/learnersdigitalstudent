package services

import (
	"fmt"
	"os"
	"task_manager_sys/internal/utils"
)

type FileProcessingTask struct {
	FileName string
}

func (fps *FileProcessingTask) Run() error {
	fmt.Println("1 File Processing Task", fps.FileName)
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(pwd)
	val, err := os.ReadFile(fps.FileName)
	if err != nil {
		return &utils.TaskError{TaskName: "FileProcessingTask", Err: err}
	}
	fileData := string(val)
	fmt.Println("val : ", fileData)
	return nil
}
