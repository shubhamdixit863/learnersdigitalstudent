package services

import (
    "errors"
    "fmt"
)

type DataValidationTask struct {
    Data string
}

func (d *DataValidationTask) Run() error {
    if d.Data == "" {
        return &TaskError{"DataValidationTask", errors.New("data cannot be empty")}
    }
    fmt.Println("Validating data:", d.Data)
    return nil
}
