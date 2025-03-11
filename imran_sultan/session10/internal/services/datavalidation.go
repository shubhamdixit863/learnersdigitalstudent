package services

import (
	"errors"
	"regexp"
)

type Datavalidation struct {
	Name  string
	Phone string
	Email string
}

func NewDataValidation() Task {
	return &Datavalidation{
		Name:  "",
		Phone: "",
		Email: "",
	}

}
func (d Datavalidation) Run() error {

	d.Name = "imran"
	d.Phone = "9541265428"
	d.Email = "examplegmail.com"
	phone := regexp.MustCompile(`^\d{10}$`)
	email := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	if !phone.MatchString(d.Phone) {
		return errors.New("invalid phone number: must be 10 digits")
	}
	if !email.MatchString(d.Email) {
		return errors.New("invalid email address format")
	}

	return nil

}
