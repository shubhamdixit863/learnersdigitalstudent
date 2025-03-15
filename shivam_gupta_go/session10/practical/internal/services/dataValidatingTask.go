package services

import (
	"fmt"
	"regexp"
	"session10/practical/internal/utils"
)

type DataValidate struct {
	Data    string
	Pattern string
}

func (a *DataValidate) Run() error {

  
	 
	re, err := regexp.Compile(a.Pattern)
	if err != nil {
		return &utils.ValidateError{Reason:"invalid regex pattern"}
	}

	if !re.Match([]byte(a.Data)) {
		return &utils.ValidateError{Reason: "pattern not match"}
	}
	fmt.Println("data validated")
	return nil
}
