package config

type ProcessingMode string

const (
	LineFilerMode    ProcessingMode = "filter"
	WordCountMode    ProcessingMode = "wordcount"
	APICallMode      ProcessingMode = "apicall"
	RetryAPICallMode ProcessingMode = "retryapi"
)

type Config struct {
	Directory     string
	Mode          ProcessingMode
	Keyword       string
	MaxGoroutines int
	APIEndpoint   string
}

func NewConfig() *Config {
	return &Config{
		MaxGoroutines: 10,
		Mode:          LineFilerMode,
		APIEndpoint:   "https://httpbin.org/post",
	}
}
