package services

import (
	"errors"
	"log"
)

type APICallTask struct {
	apiname string
}

func (s *APICallTask) Run() error {
	data, err := isAPIValid(s.apiname)
	if err != nil {
		return err
	}
	log.Println(data, "is called")
	return nil
}
func isAPIValid(name string) (string, error) {
	if name == "" {
		return "", errors.New("data is empty")
	}
	return name, nil
}
func (s *APICallTask) ReturnName() string {
	return "APICallTask"
}
func NewAPICallTask(name string) *APICallTask {
	return &APICallTask{apiname: name}
}
