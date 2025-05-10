package utils

import "fmt"

type TaskError struct {
	TaskName string
	Err      error
}

func (e *TaskError) Error() string {
	return fmt.Sprintf("task %s failed: %v", e.TaskName, e.Err)
}

type ValidateError struct {
	Message string
}

func (e *ValidateError) Error() string {
	return fmt.Sprintf("validation error: %s", e.Message)
}
