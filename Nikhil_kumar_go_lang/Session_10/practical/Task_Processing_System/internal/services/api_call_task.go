package services

import "fmt"

type APICallTask struct {
}

func (act *APICallTask) Run() error {
	fmt.Println("API Call Task")
	return nil
}
