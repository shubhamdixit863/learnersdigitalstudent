You are given a list of integers. Your task is to process each number concurrently using Goroutines while ensuring that the main function waits for all Goroutines to finish before printing a final message.

Each Goroutine should:

1. Multiply the number by 2

2. Print the result

Use sync.WaitGroup to ensure all Goroutines complete before the program exits.

Example Input:

numbers := []int{1, 2, 3, 4, 5}

Expected Output (Order may vary due to concurrency):

Processing 1 -> 2

Processing 2 -> 4

Processing 3 -> 6

Processing 4 -> 8

Processing 5 -> 10

All goroutines completed!