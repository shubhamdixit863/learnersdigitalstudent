# Concurrent File Processor with Fan-Out & Fan-In Pattern

## 📌 Overview

This Go program processes .txt files concurrently from the data/ directory using fan-out / fan-in concurrency patterns. It reads each file using goroutines and processes lines based on a selected mode.

## ⛏️ Features

1. Scans the data/ directory for .txt files automatically.
2. Uses goroutines to read files in parallel (fan-out).
3. Processes data using channels (fan-in pattern).
4. Supports multiple modes of processing:

Line Filter → Extract lines containing a specific keyword (e.g., "error").

Word Count → Count total words across all files.

API Call → Send each line to an external API & collect responses.

Retryable API Call → Retry failed API requests (max 3 times).

## 📂 Folder Structure and Setup

1. Download the zip folder from the location
2. Extract the folder to find three folders namely {"cmd"}, {"internal"} and {"data"}
3. Inside the {"cmd"} we have the main entry point of the application {"main.go"}
4. Inside the {"internal"} we furthur have four folder namely {"filehandler"}, {"services"}, {"processor"}, {"utils"}
5. These folder contain the files according to their names
6. We also have a {"data"} folder having the data to work on
7. Run the {"fileProcessor.exe"} observe the working.

## 📤 Example Output

Total words: 40
Filtered
Error: Something went wrong in module A
Info: Processing completed successfully
Warning: Low disk space

Filtered API request received
API response: 200 OK
Error: Timeout while calling external service

API Call status Success
API Call status Success
API Call status Success

## ⚠️ Error Handeling

1. FileProcessingSignal = "Processing files in directory:"
2. FileProcessingError = "Error processing files:"
3. ModeProcessingError = "Error processing mode:"
4. FileReadingError = "Error reading file:"
5. InvalidModeError = "Invalid mode"
6. APIFailed = "Failed"
