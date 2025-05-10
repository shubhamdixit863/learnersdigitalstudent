// File: config/config.go
package config

import (
	"encoding/json"
	"os"
)

// Config represents the application configuration
type Config struct {
	StartUrls             []string `json:"start_urls"`
	CrawlDepth            int      `json:"crawl_depth"`
	MaxConcurrentRequests int      `json:"max_concurrent_requests"`
	IndexFilePath         string   `json:"index_file_path"`
	ChannelBufferSize     int      `json:"channel_buffer_size"`
	DBConnectionString    string   `json:"db_connection_string"`
	DBName                string   `json:"db_name"`
}

// LoadConfig loads configuration from a JSON file
func LoadConfig(filepath string) (*Config, error) {
	// Default configuration
	config := &Config{
		StartUrls:             []string{"https://golang.org"},
		CrawlDepth:            2,
		MaxConcurrentRequests: 10,
		IndexFilePath:         "index.json",
		ChannelBufferSize:     100,
		DBConnectionString:    "localhost:27017",
		DBName:                "searchengine",
	}

	// Try to open the config file
	file, err := os.Open(filepath)
	if err != nil {
		// If file doesn't exist, create it with default values
		if os.IsNotExist(err) {
			defaultConfig, _ := json.MarshalIndent(config, "", "  ")
			err = os.WriteFile(filepath, defaultConfig, 0644)
			if err != nil {
				return nil, err
			}
			return config, nil
		}
		return nil, err
	}
	defer file.Close()

	// Decode the configuration
	decoder := json.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
