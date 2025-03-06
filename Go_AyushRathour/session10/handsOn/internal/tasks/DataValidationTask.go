package tasks

import (
	"fmt"
	"regexp"
)

type DataValidationTask struct {
	Email   string
	PhoneNo string
}

func NewDataValidationTask(email string, phone string) *DataValidationTask {
	return &DataValidationTask{
		Email:   email,
		PhoneNo: phone,
	}
}

func (task *DataValidationTask) Run() error {
	emailRegex := `^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`
	re := regexp.MustCompile(emailRegex)
	if re.MatchString(task.Email) {
		fmt.Printf("%s is valid\n", task.Email)
	} else {
		fmt.Printf("%s is invalid\n", task.Email)
	}

	re1 := regexp.MustCompile(`^\+?[0-9]{10,15}$`)
	if re1.MatchString(task.PhoneNo) {
		fmt.Printf("%s is valid\n", task.PhoneNo)
	} else {
		fmt.Printf("%s is invalid\n", task.PhoneNo)
	}

	return nil
}
