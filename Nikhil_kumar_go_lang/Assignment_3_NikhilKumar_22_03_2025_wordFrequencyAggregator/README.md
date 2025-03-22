# Word Frequency Analyzer

## Description

A high-performance concurrent word frequency analyzer implemented in Go. This tool processes text files to count and analyze word frequencies, leveraging Go's concurrency features for efficient processing of large text files.

## Features

- Concurrent text file processing using goroutines
- Chunk-based text processing for optimized memory usage
- Thread-safe operations with mutexes for result aggregation
- Displays word frequencies in descending order
- Shows total unique word count
- Efficient handling of large text files
- Clean separation of concerns with modular design

## Requirements

- Text files for analysis

Run the application by providing one or more text files as arguments:

```
./word-analyzer file1.txt
```

You can also use wildcard patterns to process multiple files:

```
./word-analyzer *.txt
```
