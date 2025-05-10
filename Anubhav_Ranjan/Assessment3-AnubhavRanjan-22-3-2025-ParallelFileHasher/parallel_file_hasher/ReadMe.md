# Parallel File Hasher

A high-performance utility that computes SHA-256 hashes of files in a directory concurrently using Go's goroutines and channels.

## Overview

This tool traverses a specified directory, computes SHA-256 hashes for each file using concurrent processing, and outputs the results in a `filename: hash` format. It's built with best practices in Go development including proper error handling, constants, SOLID principles, and a clean modular structure.

## Features

- **Concurrent Processing**: Uses a worker pool of goroutines to process files in parallel
- **Configurable Workers**: Customize the number of concurrent workers 
- **Robust Error Handling**: Comprehensive error management with descriptive messages
- **Extensible Design**: Easily add new hashing algorithms through the Strategy Pattern


## Design Principles & Best Practices

### Design Patterns

1. **Factory Method Pattern**
   - `NewFileWalker` and `NewHasherService` create and return instances with proper initialization
   - Encapsulates object creation logic

2. **Strategy Pattern**
   - `HashStrategy` interface allows swapping different hashing algorithms
   - `SHA256Strategy` implements the default hashing strategy
   - Additional strategies can be added without modifying existing code

### SOLID Principles

1. **Single Responsibility Principle**
   - Each package has a clearly defined purpose
   - `filewalker` handles directory traversal
   - `hasher` manages file hashing logic
   - `models` defines shared data structures

2. **Open/Closed Principle**
   - Code is open for extension but closed for modification
   - New hashing strategies can be added without changing existing code

3. **Dependency Inversion Principle**
   - High-level modules depend on abstractions, not implementations
   - `HasherService` depends on the `HashStrategy` interface, not concrete implementations

### Error Handling

- Error messages defined as constants to ensure consistency
- Context-rich error messages using `fmt.Errorf` with wrapping
- Proper error propagation up the call stack
- Appropriate error logging with clear messages

### Concurrency Management

- Worker pool pattern limits resource usage
- Proper channel closing to prevent goroutine leaks
- Synchronization with `sync.WaitGroup` ensures all goroutines complete
- Buffered channels to manage backpressure

### Code Organization

- Modular structure following standard Go project layout
- Clear separation of concerns between packages
- Command-line flag parsing for configuration
- Clean main function using the "functional main" pattern

