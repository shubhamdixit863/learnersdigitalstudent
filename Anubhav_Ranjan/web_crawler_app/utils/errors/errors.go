package errors

import (
	"fmt"
	"runtime"
)

// AppError represents an application error with context
type AppError struct {
	Err      error
	Message  string
	Function string
	File     string
	Line     int
}

// New creates a new application error
func New(err error, message string) *AppError {
	if err == nil {
		err = fmt.Errorf(message)
	}

	// Get caller information
	pc, file, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc).Name()

	return &AppError{
		Err:      err,
		Message:  message,
		Function: fn,
		File:     file,
		Line:     line,
	}
}

// Error implements the error interface
func (e *AppError) Error() string {
	return fmt.Sprintf("%s: %v (at %s:%d in %s)", e.Message, e.Err, e.File, e.Line, e.Function)
}

// Unwrap returns the underlying error
func (e *AppError) Unwrap() error {
	return e.Err
}

// IsType checks if the error is of a specific type
func IsType(err error, target error) bool {
	return err != nil && err.Error() == target.Error()
}

// WithMessage adds a message to an existing error
func WithMessage(err error, message string) *AppError {
	if appErr, ok := err.(*AppError); ok {
		return &AppError{
			Err:      appErr.Err,
			Message:  message,
			Function: appErr.Function,
			File:     appErr.File,
			Line:     appErr.Line,
		}
	}

	return New(err, message)
}

// HandleError handles an error with a custom handler
func HandleError(err error, handler func(error)) {
	if err != nil {
		handler(err)
	}
}

// MustNoError panics if the error is not nil
func MustNoError(err error) {
	if err != nil {
		panic(err)
	}
}

// RecoverWithHandler recovers from a panic and calls the handler
func RecoverWithHandler(handler func(interface{})) {
	if r := recover(); r != nil {
		handler(r)
	}
}

// SafeExecution executes a function and recovers from panics
func SafeExecution(fn func() error) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic recovered: %v", r)
		}
	}()

	return fn()
}
