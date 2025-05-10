package utils

import "fmt"

type TaskError struct {
	TaskName string
	Err      error
}

func (t *TaskError) Error() string {
	return fmt.Sprintf("task %s failed: %v", t.TaskName, t.Err)
}


type ValidateError struct{
	Reason string
}

func (t *ValidateError) Error()string{
	return fmt.Sprintf("validation error: %s", t.Reason)
}