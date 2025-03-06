package errors

import (
	"errors"
)

type TaskError struct {
	Message string
}

func (e *TaskError) Error() string {
	return e.Message
}

type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

func isValidationError(err error) bool {
	var ve *ValidationError
	return errors.As(err, &ve)
}

// create a add task method that takes interface Tas
