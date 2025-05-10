package services

import (
	"fmt"
	"practical/internal/model"
	"sync"
	"time"
)

func ProcessOrder(order *model.Order, inventory map[string]int, mu *sync.Mutex, wg *sync.WaitGroup, orderQueue chan model.Order) {
	defer wg.Done()
	mu.Lock()

	if (inventory)[order.Name] > 0 {
		fmt.Println("processing Order", order.Name, "ID:", order.ID)
		order.Status = "Processing"
		orderQueue <- *order
		(inventory)[order.Name]--
	} else {
		fmt.Println("Out of stock", order.Name, "ID:", order.ID)
		order.Status = "Rejected"
		orderQueue <- *order
		mu.Unlock()
		return
	}
	mu.Unlock()

	time.Sleep(1 * time.Second)
	order.Status = "Shipped"
	orderQueue <- *order
	fmt.Println("Shipped", order.Name, "ID:", order.ID)

	time.Sleep(2 * time.Second)
	order.Status = "Completed"
	orderQueue <- *order
	fmt.Println("Completed", order.Name, "ID:", order.ID)

}
