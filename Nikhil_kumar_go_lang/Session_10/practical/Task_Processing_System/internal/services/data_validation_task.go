package services

import (
	"fmt"
	"regexp"
	"task_manager_sys/internal/utils"
)

type DataValidationTask struct {
	Email         string
	Phone         string
	Age           string
	Pattern_Email string
	Pattern_Phone string
	Pattern_Age   string
}

func (dvt *DataValidationTask) Run() error {
	fmt.Println("Data Validation Task")

	email_re, err1 := regexp.Compile(dvt.Pattern_Email)
	// fmt.Println(email_re, err1, dvt.Email)
	if err1 != nil {
		return &utils.TaskError{TaskName: "DataValidationTask", Err: err1}
	}

	age_re, err2 := regexp.Compile(dvt.Pattern_Age)
	if err2 != nil {
		return &utils.TaskError{TaskName: "DataValidationTask", Err: err2}
	}

	phone_re, err3 := regexp.Compile(dvt.Pattern_Phone)
	if err3 != nil {
		return &utils.TaskError{TaskName: "DataValidationTask", Err: err3}
	}

	// Validate email
	if !email_re.MatchString(dvt.Email) {
		return &utils.ValidateError{Message: "Invalid Email"}
	}

	// Validate age
	if !age_re.MatchString(dvt.Age) {
		return &utils.ValidateError{Message: "Invalid Age"}
	}

	// Validate phone
	if !phone_re.MatchString(dvt.Phone) {
		return &utils.ValidateError{Message: "Invalid Phone"}
	}

	fmt.Println("Email:", dvt.Email, "Phone:", dvt.Phone, "Age:", dvt.Age)
	return nil
}
