package fileprocessor

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"concurrent_file_processor/internal/config"
	"concurrent_file_processor/internal/errors"
)

type ProcessingStrategy interface {
	Process(line string) (string, error)
}

type LineFilterStrategy struct {
	Keyword string
}

func (s *LineFilterStrategy) Process(line string) (string, error) {
	if strings.Contains(line, s.Keyword) {
		return line, nil
	}
	return "", nil
}

type WordCountStrategy struct {
	TotalWords int
}

func (s *WordCountStrategy) Process(line string) (string, error) {
	words := strings.Fields(line)
	s.TotalWords += len(words)
	return fmt.Sprintf("Total words: %d", s.TotalWords), nil
}

type APICallStrategy struct {
	Endpoint string
}

func (s *APICallStrategy) Process(line string) (string, error) {
	resp, err := http.Post(s.Endpoint, "application/json",
		bytes.NewBufferString(fmt.Sprintf(`{"line": "%s"}`, line)))
	if err != nil {
		return "", fmt.Errorf(errors.ErrAPICallFailed, err.Error())
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	return fmt.Sprintf("Status: %d, Response: %s", resp.StatusCode, string(body)), nil
}

type RetryAPICallStrategy struct {
	Endpoint   string
	MaxRetries int
}

func (s *RetryAPICallStrategy) Process(line string) (string, error) {
	for attempt := 0; attempt < s.MaxRetries; attempt++ {
		resp, err := http.Post(s.Endpoint, "application/json",
			bytes.NewBufferString(fmt.Sprintf(`{"line": "%s"}`, line)))

		if err == nil && resp.StatusCode < 300 {
			defer resp.Body.Close()
			body, _ := io.ReadAll(resp.Body)
			return fmt.Sprintf("Status: %d, Response: %s", resp.StatusCode, string(body)), nil
		}

		time.Sleep(time.Duration(1<<uint(attempt)) * time.Second)
	}

	return "", fmt.Errorf(errors.ErrAPICallFailed)
}

func StrategyFactory(cfg *config.Config) (ProcessingStrategy, error) {
	switch cfg.Mode {
	case config.LineFilerMode:
		return &LineFilterStrategy{Keyword: cfg.Keyword}, nil
	case config.WordCountMode:
		return &WordCountStrategy{}, nil
	case config.APICallMode:
		return &APICallStrategy{Endpoint: cfg.APIEndpoint}, nil
	case config.RetryAPICallMode:
		return &RetryAPICallStrategy{
			Endpoint:   cfg.APIEndpoint,
			MaxRetries: 3,
		}, nil
	default:
		return nil, fmt.Errorf(errors.ErrUnsupportedProcessingMode)
	}
}
