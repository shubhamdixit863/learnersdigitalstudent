package utils

import (
    "errors"
    "fmt"
    "taskprocess/internal/services" 
)

func IsTaskError(err error) bool {
    var taskErr *services.TaskError 
    return errors.As(err, &taskErr)
}

func WrapError(taskName string, err error) error {
    return fmt.Errorf("Task %s: %w", taskName, err)
}
