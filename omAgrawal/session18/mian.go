// package main
//
// import (
//
//	"fmt"
//	"math/rand"
//
// )
//
// // fan in pattern along with fan out
//
//	func generator(genCh chan<- int, count int) {
//		// it will generate random numbers
//		for range count {
//			rn := rand.Int()
//			genCh <- rn
//		}
//
//		close(genCh)
//	}
//
//	func worker(in chan int) <-chan int {
//		out := make(chan int)
//		go func() {
//			for v := range in {
//				sq := v * v
//				out <- sq
//			}
//			close(out)
//		}()
//
//		return out
//	}
//
//	func merge(resultCh chan int, n ...<-chan int) {
//		for _, v := range n {
//			for d := range v {
//				resultCh <- d
//
//			}
//		}
//		close(resultCh)
//	}
//
// const NUM_WORKERS = 5
// const COUNT = 1000
//
//	func main() {
//		genCh := make(chan int)
//		resultCh := make(chan int)
//		sliceOfChannels := make([]<-chan int, NUM_WORKERS)
//		go generator(genCh, COUNT)
//		for range NUM_WORKERS { // Workers will be spawned
//			outCh := worker(genCh)
//			sliceOfChannels = append(sliceOfChannels, outCh)
//		}
//		go func() {
//			merge(resultCh, sliceOfChannels...)
//		}()
//
//		for v := range resultCh {
//			fmt.Println(v)
//
//		}
//
// }
package main

import (
	"fmt"
	"math/rand"
)

// fan-in and fan-out pattern

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
			out <- v * v
		}
		close(out)
	}()
	return out
}

func merge(resultCh chan int, channels ...<-chan int) {
	for _, ch := range channels {
		for v := range ch {
			resultCh <- v
		}
	}
	close(resultCh)
}

const NUM_WORKERS = 5
const COUNT = 1000

func main() {
	genCh := make(chan int)
	resultCh := make(chan int)
	sliceOfChannels := make([]<-chan int, NUM_WORKERS)

	go generator(genCh, COUNT)

	// Spawn workers
	for i := 0; i < NUM_WORKERS; i++ {
		sliceOfChannels[i] = worker(genCh)
	}

	go merge(resultCh, sliceOfChannels...)

	// Print results
	for v := range resultCh {
		fmt.Println(v)
	}
}
