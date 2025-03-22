# Parallel File Hasher

A Go-based tool to compute SHA-256 hashes of all files in a directory concurrently. This tool uses goroutines to process files in parallel, making it efficient for directories with a large number of files.

---

## Features
- **Concurrent Hashing**: Uses goroutines to compute file hashes in parallel.
- **SHA-256 Hashing**: Computes the SHA-256 hash for each file.
- **Directory Scanning**: Recursively scans the specified directory for files.
- **Structured Output**: Prints the results in the format `filename: hash`.

---

## How It Works 
1. **Directory Scanning**:
   The tool walks through the specified directory and collects the paths of all files.
2. **Concurrent Hashing**:
   For each file, a goroutine is spawned to compute its SHA-256 hash.
   Results are sent back through a channel.
3. **Output**:
   The tool prints the filename and its corresponding hash.




