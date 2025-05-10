package errors

import "fmt"

type TaskError struct {
	TaskName string
	Err      error
}

func NewTaskError(taskname string, err error) *TaskError {
	return &TaskError{TaskName: taskname, Err: err}
}

func (e *TaskError) Error() string {
	return fmt.Sprintf("Task '%s' failed: %v", e.TaskName, e.Err)
}

func (e *TaskError) Unwrap() error {
	return e.Err
}
