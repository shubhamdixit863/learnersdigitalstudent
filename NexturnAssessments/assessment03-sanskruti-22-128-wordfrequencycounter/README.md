# Word Frequency Aggregator

## Overview
The Word Frequency Aggregator is a Go-based application that reads a large text file (`big.txt`), splits it into chunks, and calculates the frequency of words using goroutines for parallel processing. The results are then aggregated and displayed.

## Features
- Efficiently reads large text files in chunks.
- Uses goroutines for parallel processing.
- Calculates word frequencies using unbuffered channels.
- Aggregates and displays the top frequent words.

## Error Handling
- Logs errors during file reading, chunking, or frequency calculation.
- Gracefully handles panics and closed channels.

## Code Explantion of each file

1. cmd/main.go
- This is the entry point of the application.
- It initializes the channels for data processing (`ch`, `resultCh`, and doneCh).
- Launches goroutines for:
  - *Reading the file* using `services.ReadFile()`.
  - *Processing the data* using `services.WordScanner()` and `services.CalculateFrequency()`.
  - *Aggregating results* using `services.AggregateResults()`.
- Prints the final word frequency results

2. internal/services/aggregator.go
- Contains the logic for aggregating word frequency results.
- It listens to the results from the `resultCh` channel and merges frequency maps into a final result map.
- Sends a signal through the `doneCh` channel when aggregation is complete.


3.internal/sefvices/chunker.go
- Responsible for splitting the text data into words using a scanner.
- Uses `bufio.Scanner` to read data chunks and apply a custom split function to extract words.
- Handles data cleaning by removing unwanted characters or symbols.


4.internal/services/freqcounter.go
- Contains the logic to calculate word frequencies.
- It accepts a slice of words and generates a map where keys are words and values are their frequency count.
- Efficiently updates the frequency map using a simple loop.

5. internal/services/reader.go
- Handles reading the large file in chunks using `os.Open()` and `io.ReadAtLeast()`.
- Uses a channel to send the read chunks of data to the processing goroutines.
- Ensures the channel is closed once all chunks are read.

---

6. internal/models/constants.go
- Contains constants and configuration values used throughout the application.
- Variables include:
  - `FilePath` - Path to the input text file (`big.txt`).
  - `ChunkSize` - Size of the chunks to read (in bytes).
  - `NumWorkers` - Number of goroutines used for parallel processing.
  - `TopWordCount` - Number of top frequent words to display.

---

7. big.txt
- This is the sample input file for the word frequency calculation.
- It should contain large amounts of text for analysis.

8. go.mod
- These are Go module files.
- `go.mod` tracks dependencies and module versions.

## Conclusion
- Each component is well-separated using Go’s package system, following clean coding practices.
- The application uses channels for concurrency and ensures proper error handling and resource management.
- Results are displayed efficiently using the final frequency map after aggregation.

