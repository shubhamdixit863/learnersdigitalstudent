# 📂 Real-Time Directory Monitor (Simulated)

## Overview

This Go application simulates real-time file processing by monitoring a directory, feeding file paths through a channel, and processing files concurrently using goroutines. Each file is read, and its line count is printed.

## Features

1. Concurrent file processing
2. Concurrent file processing using goroutines and channels
3. Buffered channel for controlled processing
4. Dynamic file discovery from a directory
5. Graceful error handling

## Installation and Setup

1. Unzip the folder downloaded
2. Find a folder with the folder structure consisting of {"cmd"}, {"internal"}
3. Inside the {"cmd"} we have the main entry point of the application {"main.go"}
4. Inside the {"internal"} we furthur have four folder namely {"monitor"}, {"services"}, {"files"}, {"utils"}, {"data"}
5. These folder contain the files according to their names

## How it works

1. The program scans the ./internal/files as well as ./internal/data directory for text files.
2. Each file path is sent through a buffered channel.
3. Multiple goroutines process files concurrently.
4. Each file is read line by line, and the total line count is printed.
5. The program waits until all files are processed before exiting.

## Example output

File: ./internal/files/file1.txt, Lines: 1
File: ./internal/files/file2.txt, Lines: 2
File: ./internal/files/file3.txt, Lines: 1
File: ./internal/files/book.txt, Lines: 4
File: ./internal/data/file1.txt, Lines: 3
File: ./internal/data/file2.txt, Lines: 4
File: ./internal/data/file3.txt, Lines: 1
2025/03/22 12:44:17 Error processing file ./internal/files/file4.txt: failed to open file: open ./internal/files/file4.txt: The system cannot find the file specified.
All files processed

## Error Handling

1. FileProcessErr = "Error processing file %s: %v"
2. FileOpenErr = "failed to open file: %w"
3. FileReadErr = "error reading file: %w"
