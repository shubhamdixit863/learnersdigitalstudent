package services

import (
	"fmt"
	"task/internal/utils"
)

type DataValidation struct {
	Data string
}

func (d *DataValidation) Run() error {
	fmt.Println("Validating data", d.Data)

	if (d.Data == "" || d.Data == "Throw error"){
		return &utils.ValidationError{Message: "Data is empty"}
	}

	return nil
}

