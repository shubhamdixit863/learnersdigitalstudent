package services

import (
	"fmt"
	"practical/internal/errorss"
	"regexp"
)

type DataValidation struct {
	Email string
}

func NewDataValidation(email string) Task {
	return &DataValidation{Email: email}
}

func (d *DataValidation) Run() error {
	var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(d.Email) {
		return &errorss.ValidationError{Msg: "Enter a valid email."}
	}
	fmt.Printf("Email %s is Valid.\n", d.Email)
	return nil
}
