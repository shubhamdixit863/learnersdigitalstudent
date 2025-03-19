package services

import (
	"fmt"
	"session10/problems/internal/utils"
)

type APICallTask struct {
	Link string
}

func NewAPICallTask(link string) *APICallTask {
	return &APICallTask{
		Link: link,
	}
}

func (f APICallTask) Run() error {
	defer fmt.Println("APICallTask: Cleanup")
	fmt.Println("Calling API Task...")

	if f.Link == "" {
		return fmt.Errorf("%w: Link is empty", utils.ValidationError)
	}

	fmt.Printf("Calling API Call Link: %s\n", f.Link)
	return nil
}
