package tasks

import "fmt"

type APICallTask struct {
	Endpoint string
}

func NewAPICallTask(api string) *APICallTask {
	return &APICallTask{
		Endpoint: api,
	}
}

func (task *APICallTask) Run() error {
	fmt.Println("Calling API:", task.Endpoint)
	return fmt.Errorf("API call failed for %s", task.Endpoint)
}
