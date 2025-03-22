## Concurrent Csv Aggregator

In this code I have created a CSV Aggregator which takes the csv file as input and splits the file in N chunks
makes their respective Go Routine and finds the total amount spend by per user.


 ## 1. Read & Split CSV File (reader package)

The CSV file is split into chunks of fixed size (chunkSize).

The chunks are read into memory and passed to worker goroutines for processing.

 ## 2. Process Chunks Concurrently (processor package)

Each chunk is processed by a separate goroutine.

A map of user IDs and their total spend (map[string]float64) is generated and sent through a buffered channel.

 ## 2. Merge Results (utils package)

   The main() function receives maps from the channel and merges them in a thread-safe manner using a mutex.


