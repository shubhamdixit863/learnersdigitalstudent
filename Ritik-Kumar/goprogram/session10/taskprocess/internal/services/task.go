package services

import "fmt"

type Task interface {
	Run() error
}

type TaskError struct {
	TaskName string
	Err      error
}

func (e *TaskError) Error() string {
	return fmt.Sprintf("Task %s failed: %v", e.TaskName, e.Err)
}
