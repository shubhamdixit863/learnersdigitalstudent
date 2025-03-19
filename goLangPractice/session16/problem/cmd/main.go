package main

import (
	"session16/problem/internal/services"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	orderQueue := make(chan services.Order, 2)
	orderStatus := make(chan []string, 2)

	order1 := services.CreateNewOrder(1, map[string]int{"Mobile": 2, "Laptop": 1})
	order2 := services.CreateNewOrder(2, map[string]int{"Washing Machine": 1, "Charger": 1, "Cable": 3})
	order3 := services.CreateNewOrder(3, map[string]int{"Jeans": 2, "Shirt": 4})
	order4 := services.CreateNewOrder(4, map[string]int{"Jeans": 3, "Shirt": 3})

	orders := []services.Order{*order1, *order2, *order3, *order4}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for order := range orderQueue {
			services.UpdateInventory(order, orderStatus, &mu)
		}

		close(orderStatus)

	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for status := range orderStatus {
			services.UpdateStatus(status)
		}
	}()

	for _, order := range orders {
		wg.Add(1)
		go func(o services.Order) {
			defer wg.Done()
			o.PlaceOrder(orderQueue)
		}(order)
	}

	wg.Wait()
	close(orderQueue)

}

//package main
//
//import (
//	"session16/problem/internal/services"
//	"sync"
//)
//
//func main() {
//	var wg sync.WaitGroup      // For processing and status updates
//	var orderWg sync.WaitGroup // For order placements
//	var mu sync.Mutex
//	orderQueue := make(chan services.Order)
//	orderStatus := make(chan []string)
//
//	order1 := services.CreateNewOrder(1, map[string]int{"Mobile": 2, "Laptop": 1})
//	order2 := services.CreateNewOrder(2, map[string]int{"Washing Machine": 1, "Charger": 1, "Cable": 3})
//	order3 := services.CreateNewOrder(3, map[string]int{"Jeans": 2, "Shirt": 4})
//
//	orders := []services.Order{*order1, *order2, *order3}
//
//	// Start order processing goroutine
//	wg.Add(1)
//	go func() {
//		defer wg.Done()
//		for order := range orderQueue {
//			services.UpdateInventory(order, orderStatus, &mu)
//		}
//		close(orderStatus)
//	}()
//
//	// Start status update goroutine
//	wg.Add(1)
//	go func() {
//		defer wg.Done()
//		for status := range orderStatus {
//			services.UpdateStatus(status)
//		}
//	}()
//
//	// Place orders
//	orderWg.Add(len(orders))
//	for _, order := range orders {
//		go func(o services.Order) {
//			defer orderWg.Done()
//			o.PlaceOrder(orderQueue)
//		}(order)
//	}
//
//	// Wait for all orders to be placed and close the queue
//	orderWg.Wait()
//	close(orderQueue)
//
//	// Wait for processing and status updates to complete
//	wg.Wait()
//}
