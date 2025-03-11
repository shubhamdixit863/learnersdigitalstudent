package services

import (
	"net/http"
)

type ApiCall struct {
	url string
}

func NewApiCall() Task {
	return &ApiCall{
		url: "",
	}
}
func (a ApiCall) Run() error {
	url := "https://www.google.com/"

	_, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	return nil
}
