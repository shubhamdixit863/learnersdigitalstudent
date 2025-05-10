package errorss

import "fmt"

type TaskError struct {
	Msg string
}

func (e *TaskError) Error() string {
	return fmt.Sprintf("TaskError: %s", e.Msg)
}

type ValidationError struct {
	Msg string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("ValidationError: %s", e.Msg)
}
