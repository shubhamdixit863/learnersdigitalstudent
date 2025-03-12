package services

import (
    "errors"
    "fmt"
)

type APICallTask struct {
    URL string
}

func (a *APICallTask) Run() error {
    fmt.Println("Calling API:", a.URL)
    return &TaskError{"APICallTask", errors.New("API request failed")}
}
