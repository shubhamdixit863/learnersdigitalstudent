package utils

import "fmt"

type Task interface {
	Run() error
}

type TaskError struct {
	Message string
}

func (e *TaskError) Error() string {
	return fmt.Sprintf("Task error: %s", e.Message)
}

type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("Validation error: %s", e.Message)
}
