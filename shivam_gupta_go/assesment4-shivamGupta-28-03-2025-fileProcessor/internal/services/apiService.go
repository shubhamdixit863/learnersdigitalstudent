package services

import (
	"bytes"
	"encoding/json"
	"fileProcessor/internal/models"
	"fmt"
	"io"
	// "log"
	"net/http"
	"net/url"
	"time"
)

const (
	apiBaseURL  = "https://httpbin.org"
	apiStatus1  = "POST: %s\nGET: %s"
	post1       = "/post"
	post        = "POST"
	err1        = "POST request creation failed: %v"
	apiStatus2  = "POST failed: request creation error"
	contentType = "Content-Type"
	text        = "text/plain"
	err2 = "POST request failed: %v"
	postFailed1 ="POST failed: request execution error"
	postFailed2 = "POST failed: response read error"
	sprint1 = "%s/get?data=%s"
	get = "GET"
	err3 = "GET request creation failed: %v"
	apiStatus3 = "GET failed: request creation error"
	args = "args"
	data = "data"
  dataRecieved = "GET Request Data Received:"
	status1 = "GET data not found in response"
	status2 = "GET failed: data not in response"
)

func APICall(line models.LineWithMetadata) models.ProcessResult {

	postResult := makePostRequest(line.Line)
	if postResult.Error != nil {
		return postResult
	}

	getResult := makeGetRequest(line.Line)

	return models.ProcessResult{
		FileName:   line.FileName,
		LineNum:    line.LineNum,
		Line:       line.Line,
		APISuccess: postResult.APISuccess ,
		APIStatus:  fmt.Sprintf(apiStatus1, postResult.APIStatus, getResult.APIStatus),
	}
}

func makePostRequest(data string) models.ProcessResult {
	url := apiBaseURL + post1
	body := bytes.NewBufferString(data)

	req, err := http.NewRequest(post, url, body)
	if err != nil {
		return models.ProcessResult{
			Error:     fmt.Errorf(err1, err),
			APIStatus: apiStatus2,
		}
	}

	req.Header.Set(contentType, text)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return models.ProcessResult{
			Error:     fmt.Errorf(err2, err),
			APIStatus: postFailed1,
		}
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.ProcessResult{
			Error:     fmt.Errorf(err2, err),
			APIStatus: postFailed2,
		}
	}

	return models.ProcessResult{
		APISuccess: resp.StatusCode == http.StatusOK,
		APIStatus:  string(respBody),
	}
}

func makeGetRequest(data string) models.ProcessResult {
	encodedData := url.QueryEscape(data)
	url := fmt.Sprintf(sprint1, apiBaseURL, encodedData)

	req, err := http.NewRequest(get, url, nil)
	if err != nil {
		return models.ProcessResult{
			Error:     fmt.Errorf(err3, err),
			APIStatus: apiStatus3,
		}
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return models.ProcessResult{
			Error:     fmt.Errorf(err3, err),
			APIStatus: apiStatus3,
		}
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.ProcessResult{
			Error:     fmt.Errorf(err3, err),
			APIStatus: apiStatus3,
		}
	}

	
	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return models.ProcessResult{
			Error:     fmt.Errorf(err3, err),
			APIStatus: apiStatus3,
		}
	}

	
	if args, ok := result[args].(map[string]interface{}); ok {
		if receivedData, ok := args[data].(string); ok {
			fmt.Println(dataRecieved, receivedData)
		} else {
			return models.ProcessResult{
				Error:     fmt.Errorf(err3 ,ok),
				APIStatus: status2,
			}
		}
	}

	return models.ProcessResult{
		APISuccess: resp.StatusCode == http.StatusOK,
		APIStatus:  string(respBody),
	}
}

func RetryableAPICall(line models.LineWithMetadata) models.ProcessResult {
	var err error
	var resp *http.Response

	for i := 0; i < 3; i++ {
		resp, err = http.Post(apiBaseURL, text, bytes.NewBuffer([]byte(line.Line)))
		if err == nil {
			defer resp.Body.Close()
			body, _ := io.ReadAll(resp.Body)
			return models.ProcessResult{
				FileName:   line.FileName,
				LineNum:    line.LineNum,
				Line:       line.Line,
				APISuccess: true,
				APIStatus:  string(body),
			}
		}
		time.Sleep(time.Duration(i+1) * time.Second)
	}

	return models.ProcessResult{
		FileName:   line.FileName,
		LineNum:    line.LineNum,
		Line:       line.Line,
		APISuccess: false,
		Error:      fmt.Errorf(err3, err),
	}
}
