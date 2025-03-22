// package services
//
// import (
//
//	"fmt"
//	"math/rand"
//	"session16/storage"
//	"sync"
//
// )
//
//	func CreateMultipleorder(wg *sync.WaitGroup, ch chan *storage.Order) {
//		for i := 0; i < 4; i++ {
//			qnt := rand.Intn(15)
//			wg.Add(1)
//			go CreateOrder(i, qnt, wg, ch)
//		}
//		wg.Add(1)
//		go ProcessOrders(wg, ch) // Start a goroutine to process orders
//
//		wg.Wait()
//
// }
//
//	func CreateOrder(id int, qnty int, wg *sync.WaitGroup, ch chan *storage.Order) {
//		defer wg.Done() // Ensure it's called to decrement the WaitGroup counter
//		order := storage.Neworder(id, qnty, "Processing")
//		ch <- order // This won't block now
//		fmt.Println("Order created:", *order)
//	}
//
//	func ProcessOrders(wg *sync.WaitGroup, ch chan *storage.Order) {
//		for order := range ch {
//			UpdateInventory(order)
//		}
//		wg.Done()
//	}
//
//	func UpdateInventory(order *storage.Order) {
//		qnt := order.Qunty
//		storage.Mu.Lock()
//		defer storage.Mu.Unlock()
//		storage.QuantityWarehouse -= qnt
//		storage.OrderList = append(storage.OrderList, order)
//		order.Status = "Shipped"
//		fmt.Println("Order shipped:", *order)
//		fmt.Println(storage.QuantityWarehouse)
//
// }
// services/order_service.go
package services

import (
	"fmt"
	"math/rand"
	"session16/storage"
	"sync"
)

func CreateMultipleorder(wg *sync.WaitGroup, ch chan *storage.Order) {
	// Create multiple orders
	orderWg := sync.WaitGroup{}
	for i := 0; i < 11; i++ {
		qnt := rand.Intn(15)
		orderWg.Add(1)
		go func(id int, quantity int) {
			defer orderWg.Done()
			CreateOrder(id, quantity, wg, ch)
		}(i, qnt)
	}

	// Wait for all order creation to complete
	orderWg.Wait()

	close(ch)

	// Start a goroutine to process orders
	wg.Add(1)
	go ProcessOrders(wg, ch)
}

func CreateOrder(id int, qnty int, wg *sync.WaitGroup, ch chan *storage.Order) {
	order := storage.Neworder(id, qnty, "Processing")
	ch <- order
	fmt.Println("Order created:", *order)
}

func ProcessOrders(wg *sync.WaitGroup, ch chan *storage.Order) {
	defer wg.Done()

	// Process orders until channel is closed
	for order := range ch {
		UpdateInventory(order)
	}

	// After processing all orders, print the final state
	fmt.Println("All orders processed. Final inventory:", storage.QuantityWarehouse)
}

func UpdateInventory(order *storage.Order) {
	qnt := order.Qunty
	storage.Mu.Lock()
	defer storage.Mu.Unlock()
	storage.QuantityWarehouse -= qnt
	storage.OrderList = append(storage.OrderList, order)
	order.Status = "Shipped"
	fmt.Println("Order shipped:", *order)
	fmt.Println("Remaining inventory:", storage.QuantityWarehouse)
}
