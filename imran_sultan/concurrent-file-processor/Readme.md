Concurrent File Processor with Fan-Out and Fan-In Pattern

Overview

This project is a concurrent file processing system built in Go. It reads lines from multiple .txt files in a given directory and processes them concurrently using the fan-out / fan-in concurrency pattern.

Features

Accepts a directory path as input.

Scans the directory for .txt files.

Reads files concurrently using goroutines (fan-out pattern).

Sends lines to a central channel for processing (fan-in pattern).

Supports multiple processing modes:

Line Filter: Filters lines containing a specific keyword (e.g., "error").

Word Count: Counts the total number of words across all lines.

API Call: Sends each line to an external API (httpbin.org/post) and collects the response status.

Retryable API Call: Sends each line to an API with up to 3 retries in case of failure, using an exponential backoff strategy.

Prints the results based on the selected processing mode.

Usage

Prerequisites

