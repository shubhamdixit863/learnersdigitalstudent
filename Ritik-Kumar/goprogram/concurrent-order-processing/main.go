package main

import (
	"math/rand"
	"sync"
	"time"

	"concurrent-order-processing/orders"
	"concurrent-order-processing/utils"
)

func main() {
	numOrders := 20 // Number of orders to process
	orderQueue := make(chan orders.Order, numOrders)
	inventory := orders.NewInventory()

	// Set initial stock for products
	products := []string{"Product-1", "Product-2", "Product-3", "Product-4", "Product-5"}
	for _, product := range products {
		inventory.SetStock(product, 10)
	}

	// Generate random orders
	var orderList []orders.Order
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < numOrders; i++ {
		orderList = append(orderList, orders.NewOrder(
			products[rand.Intn(len(products))], // Random product
			rand.Intn(5)+1,                     // Random quantity (1 to 5)
		))
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go orders.ProcessOrders(orderQueue, inventory, &wg)

	utils.Log(" Processing Orders...")

	// Send orders to queue
	for _, order := range orderList {
		orderQueue <- order
	}
	close(orderQueue) // Close channel after sending all orders

	wg.Wait()
	utils.Log(" All orders processed. Shutting down gracefully.")
}
