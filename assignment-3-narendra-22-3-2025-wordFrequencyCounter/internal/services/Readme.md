1. main.go 
Creates an Aggregator instance to handle file reading and word frequency counting.

Opens the file safely and ensures proper error handling.

Starts two goroutines:

One reads the file (Reader).

Another processes word frequency (FrequencyCounter).

Collects and prints the final word frequencies.

Gracefully handles errors instead of crashing.



2. reader.go (File Reader)
Reads the text file line by line using bufio.Scanner.

Sends each line to chunkChan (unbuffered channel) for processing.

Handles any file reading errors 

Closes the channel when reading is done to prevent deadlocks.


3. chunker.go 
Splits a line into words using strings.Fields().

Returns a list of words for further processing.



4. frequencycounter.go 
Receives lines from chunkChan, breaks them into words using Chunker().

Counts how many times each word appears.

Sends the final word count to resultChan.

Handles cases where no words are found (warns the user).

Closes resultChan when done to avoid blocking operations.

