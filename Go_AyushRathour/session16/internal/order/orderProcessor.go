package order

import (
	"fmt"
	"math/rand"
	"session16/internal/inventory"
	"session16/internal/status"
	"sync"
	"time"
)

func ProcessOrders(wg *sync.WaitGroup) {
	defer wg.Done()
	for order := range OrderQueue {
		if inventory.UpdateInventory(order.Item, order.Quantity) {
			order.Status = "Shipped"
			status.Queue <- fmt.Sprintf("Order %d: %s", order.OrderID, order.Status)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			order.Status = "Completed"
			status.Queue <- fmt.Sprintf("Order %d: %s", order.OrderID, order.Status)
		} else {
			status.Queue <- fmt.Sprintf("Order %d: Rejected (Insufficient Stock)", order.OrderID)
		}
	}
}
