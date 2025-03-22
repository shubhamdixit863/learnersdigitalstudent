go run main.go

Expected Output
the: 4
quick: 3
brown: 1
fox: 2
jumps: 1
lazy: 3
dog: 2
is: 2
and: 1
thinking: 1
saves: 1
time: 1
coding: 1
causes: 1
problems: 1

File Reading (reader/reader.go)

Reads the input text file in chunks of lines (default: 2 lines per chunk).

Uses bufio.Scanner to efficiently scan through the file.

Sends each chunk as a slice of strings to an unbuffered channel.

2️ Parallel Processing (processor/counter.go)

Each chunk is received and processed in a separate goroutine.

The words are cleaned (punctuation removed) and counted.

Results are sent to another unbuffered channel for merging.

3️ Merging Results (utils/merger.go)

Collects word count maps from all goroutines.

Merges the results into a final frequency map.

Prints the final word count statistics.














// go func() {
	// 	for chunk := range chunks {
	// 		wg.Add(1)
	// 		go processor.CountWords(chunk, results, &wg)
	// 	}

	// 	wg.Wait()
	// 	close(results)
	// }()