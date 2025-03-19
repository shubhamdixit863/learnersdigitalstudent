package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"

	"session16/internal/inventory"
	"session16/internal/order"
	"session16/internal/shutdown"
	"session16/internal/status"
)

func main() {
	var wg sync.WaitGroup

	inventory.InitInventory()

	for i := 0; i < 5; i++ { // Start processing orders
		wg.Add(1)
		go order.ProcessOrders(&wg)
	}

	wg.Add(1) // Start status logger
	go status.StatusLogger(&wg)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		shutdown.GracefulShutdown()
	}()

	// Simulating incoming orders
	for i := 0; i < 50; i++ {
		go order.PlaceOrder("itemA", 1)
		go order.PlaceOrder("itemB", 2)
		go order.PlaceOrder("itemC", 3)
		time.Sleep(50 * time.Millisecond) // Simulate order interval
	}

	wg.Wait()
	fmt.Println("All orders processed successfully.")
}
