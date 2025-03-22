package services

import (
	"fmt"
	"task/internal/utils"
)

type APICall struct {
	Endpoint string
}

func (a *APICall) Run() error {
	fmt.Println("Running API call to", a.Endpoint)

	return &utils.TaskError{Message: "API call failed"}
}

