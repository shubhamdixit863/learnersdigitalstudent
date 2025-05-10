package tasks

import (
	"fmt"
	"regexp"
	"task_processing_system/internal/services/errors"
)

type DataValidationTask struct {
	UserName string
	Email    string
	Phone    string
}

func NewDataValidationTask(username, email, phone string) *DataValidationTask {
	return &DataValidationTask{UserName: username, Email: email, Phone: phone}
}

func (dvt *DataValidationTask) Run() error {
	nameRegex := `^[A-Za-z\s]{2,50}$`
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	phoneRegex := `^\+?[0-9]{10,15}$`

	if match, _ := regexp.MatchString(nameRegex, dvt.UserName); !match {
		return errors.NewValidationError("UserName", "Name must contain only letters and spaces, and be 2 to 50 characters long")
	}

	if match, _ := regexp.MatchString(emailRegex, dvt.Email); !match {
		return errors.NewValidationError("Email", "Invalid email format")
	}

	if match, _ := regexp.MatchString(phoneRegex, dvt.Phone); !match {
		return errors.NewValidationError("Phone", "Invalid phone number format")
	}

	fmt.Printf("Validation successful for user: %s\n", dvt.UserName)
	return nil
}

func (dvt *DataValidationTask) Name() string {
	return "DataValidationTask"
}
