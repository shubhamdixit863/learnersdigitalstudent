package services

import (
	"fmt"
	"regexp"
	"session10/problems/internal/utils"
)

type DataValidationTask struct {
	Name  string
	Email string
	Phone string
}

func NewDataValidationTask(name string, email string, phone string) *DataValidationTask {
	return &DataValidationTask{
		Name:  name,
		Email: email,
		Phone: phone,
	}
}

func (f DataValidationTask) Run() error {
	if f.Name == "" {
		return fmt.Errorf("%w: Name is empty", utils.ValidationError)
	}

	emailRegex := `^[a-zA-Z0-9._%%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	if matched, _ := regexp.MatchString(emailRegex, f.Email); !matched {
		return fmt.Errorf("%w: Invalid email format", utils.ValidationError)
	}

	phoneRegex := `^[0-9]{10}$`
	if matched, _ := regexp.MatchString(phoneRegex, f.Phone); !matched {
		return fmt.Errorf("%w: Invalid phone number", utils.ValidationError)
	}

	fmt.Println("Data validated successfully")
	return nil
}
