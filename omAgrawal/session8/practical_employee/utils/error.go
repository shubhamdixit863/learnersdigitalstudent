package utils

import "C"

type CustomError struct {
	Message  string
	Provider string
}

func (c *CustomError) Error() string {
	return c.Message + c.Provider
}

func CustomeErrorFunc(messsage string, provider string) error {

	return &CustomError{Message: messsage, Provider: provider}

}
