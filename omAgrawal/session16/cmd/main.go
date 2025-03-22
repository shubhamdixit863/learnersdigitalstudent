package main

import (
	"session16/internals/services"
	"session16/storage"
	"sync"
)

var OrderQueue = make(chan *storage.Order) // Buffered channel with size 4
var wg sync.WaitGroup

func main() {
	// Create orders and start processing
	services.CreateMultipleorder(&wg, OrderQueue)

	// Wait for all goroutines to complete

	//wg.Add(1)
	//go services.ProcessOrders(&wg, OrderQueue)
	wg.Wait()
	// Don't close the channel here anymore - it's handled in CreateMultipleorder
}
