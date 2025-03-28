package main

import (
	"fmt"
	"math/rand"
)

// Fan-in and Fan-out pattern

func generator(genCh chan<- int, count int) {
	// Generate random numbers
	for i := 0; i < count; i++ {
		rn := rand.Int()
		genCh <- rn
	}
	close(genCh)
}

func worker(in chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			sq := v * v
			out <- sq
		}
		close(out)
	}()
	return out
}

func merge(resultCh chan int, channels ...<-chan int) {
	for _, ch := range channels {
		for val := range ch {
			resultCh <- val
		}
	}
	close(resultCh)
}

const NUM_WORKERS = 5
const COUNT = 1000

func main() {
	genCh := make(chan int)
	resultCh := make(chan int)

	sliceOfChannels := make([]<-chan int, 0, NUM_WORKERS) // Correct slice initialization

	// Start the generator
	go generator(genCh, COUNT)

	// Start worker goroutines
	for i := 0; i < NUM_WORKERS; i++ {
		outCh := worker(genCh)
		sliceOfChannels = append(sliceOfChannels, outCh)
	}

	// Start merging results
	go func() {
		merge(resultCh, sliceOfChannels...)
	}()

	// Print results
	for v := range resultCh {
		fmt.Println(v)
	}
}
