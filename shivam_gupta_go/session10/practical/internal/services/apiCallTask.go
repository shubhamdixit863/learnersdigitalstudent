package services

import (
	"fmt"
)

type ApiCall struct {
	Filepath string
}

func (a *ApiCall) Run() error {
	fmt.Println("Api call.......")
	
	return nil
}

