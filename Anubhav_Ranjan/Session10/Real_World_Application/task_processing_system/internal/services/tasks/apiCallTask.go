package tasks

import (
	"fmt"
	"net/http"
	"task_processing_system/internal/services/errors"
)

type APICallTask struct {
	URL string
}

func NewAPICallTask(url string) *APICallTask {
	return &APICallTask{URL: url}
}

func (act *APICallTask) Run() error {
	resp, err := http.Get(act.URL)
	if err != nil {
		apiErr := fmt.Errorf("API call failed: %w", err)
		return errors.NewTaskError(act.Name(), apiErr)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respErr := fmt.Errorf("API call error: received status code %d", resp.StatusCode)
		return errors.NewTaskError(act.Name(), respErr)
	}

	fmt.Println("API call successful:", act.URL)
	return nil
}

func (act *APICallTask) Name() string {
	return "APICallTask"
}
