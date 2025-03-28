package services

import (
	"errors"
	"fmt"
	"goConcurrencyHandsOn/internal/utilities"
	"net/http"
	"strings"
)

func ProcessLines(lines chan string, mode int, keyword string) {
	switch mode {
	case 1:
		var count int
		fmt.Println(utilities.ModeOne, keyword)
		for line := range lines {
			if strings.Contains(line, keyword) {
				fmt.Println(line)
				count++
			}
		}
		if count == 0 {
			fmt.Println(utilities.ModeOneFail, keyword)
		}
	case 2:
		var count int
		for line := range lines {
			lineSlice := strings.Fields(line)
			count += len(lineSlice)
		}
		fmt.Println(utilities.ModeTwo, count)
	case 3:
		lineStatus := make(map[string]string)

		for line := range lines {
			status, err := sendAPIRequest(line)
			if err != nil {
				fmt.Println(utilities.ModeThreeFail, err)
			} else {
				lineStatus[line] = status
			}
		}

		fmt.Println(utilities.ModeThree)
		for k, v := range lineStatus {
			fmt.Println(k, utilities.Colon, v)
		}
	case 4:
		lineStatus := make(map[string]string)
		var maxRetries = 3

		for line := range lines {
			status, err := retryAPIRequest(line, maxRetries)
			if err != nil {
				fmt.Println(err)
			} else {
				lineStatus[line] = status
			}
		}

		fmt.Println(utilities.ModeThree)
		for k, v := range lineStatus {
			fmt.Println(k, utilities.Colon, v)
		}
	default:
		fmt.Println(utilities.DefaultError)
	}
}

func sendAPIRequest(line string) (string, error) {
	resp, err := http.Post(utilities.URL, utilities.ContentType, strings.NewReader(line))
	if err != nil {
		return utilities.EmptyString, err
	}
	defer resp.Body.Close()

	return resp.Status, nil
}

func retryAPIRequest(line string, maxRetries int) (string, error) {
	for i := 0; i < maxRetries; i++ {
		status, err := sendAPIRequest(line)
		if err == nil {
			return status, nil
		}
	}
	return "", errors.New(utilities.ModeFourFail)
}
