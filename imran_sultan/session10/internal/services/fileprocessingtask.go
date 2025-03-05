package services

import (
	"fmt"
	"os"
)

type FileProcessing struct {
	track []string
}

func NewTask() Task {

	
	return &FileProcessing{
		track: track=append(track, "im"),
	}
}

func ReadFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	return string(data), err

}
