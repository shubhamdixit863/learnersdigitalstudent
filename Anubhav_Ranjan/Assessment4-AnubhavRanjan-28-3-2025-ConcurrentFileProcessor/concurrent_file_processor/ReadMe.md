# Concurrent File Processor

## Overview

This Go application demonstrates an advanced concurrent file processing system using fan-out and fan-in concurrency patterns. The project showcases best practices in Go programming, including modular architecture, error handling, and flexible processing strategies.

## Features

- **Concurrent File Processing**: Read and process multiple text files simultaneously
- **Multiple Processing Modes**:
  - Line Filtering
  - Word Counting
  - API Calling
  - Retryable API Calling
- **Robust Error Handling**
- **Configurable Processing**
- **Extensible Design**

## Project Structure



## Key Design Patterns

### 1. Concurrency Patterns
- **Fan-Out**: Multiple goroutines read files concurrently
- **Fan-In**: Centralized channel for processing lines
- Uses `sync.WaitGroup` for goroutine management
- Context-based cancellation and timeout

### 2. Strategy Pattern
Implements different processing strategies:
- `LineFilterStrategy`: Filter lines by keyword
- `WordCountStrategy`: Count words across files
- `APICallStrategy`: Send lines to external API
- `RetryAPICallStrategy`: Retry API calls with backoff

### 3. Factory Method
`StrategyFactory` creates appropriate processing strategy based on configuration

## Configuration Options

The application supports multiple configuration parameters:

- `-dir`: Directory to process (default: current directory)
- `-mode`: Processing mode 
  - `filter`: Line filtering
  - `wordcount`: Word counting
  - `apicall`: Simple API calls
  - `retryapi`: Retryable API calls
- `-keyword`: Keyword for line filtering (default: "error")

## Usage Examples

```bash
# Line Filtering Mode
go run cmd/main.go -dir ./testdata -mode filter -keyword error

# Word Count Mode
go run cmd/main.go -dir ./testdata -mode wordcount

# API Call Mode
go run cmd/main.go -dir ./testdata -mode apicall

# Retryable API Call Mode
go run cmd/main.go -dir ./testdata -mode retryapi
```

## Error Handling

The application implements robust error handling:
- Custom error types
- Error wrapping with context
- Centralized error logging
- Concurrent error propagation

## Performance Considerations

- Buffered channels for efficiency
- Controlled goroutine count
- Context-based timeout
- Exponential backoff for API retries

## Dependencies

- Standard Go libraries

## Installation

1. Clone the repository
2. Ensure Go 1.21+ is installed
3. Run `go mod tidy` to download dependencies

## Extending the Application

To add new processing strategies:
1. Implement the `ProcessingStrategy` interface
2. Add a new case in `StrategyFactory`

## About

This project demonstrates advanced Go concurrency patterns, modular design, and best practices in software development.
