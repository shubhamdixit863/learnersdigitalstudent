package services

import (
	"fmt"
	"io"
	"net/http"
	"practical/internal/errorss"
)

type APICall struct {
	URL      string
	Response string
}

func NewAPICall(url string) Task {
	return &APICall{URL: url}
}
func (a *APICall) Run() error {
	response, err := http.Get(a.URL)

	if err != nil {
		return &errorss.TaskError{Msg: fmt.Sprintf("%v", err)}
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return &errorss.TaskError{Msg: fmt.Sprintf("%v", err)}
	}
	a.Response = string(responseData)
	fmt.Println("API Called: ")
	fmt.Println("Response is:", string(responseData))
	return nil
}
