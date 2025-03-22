package orders

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"concurrent-order-processing/utils"
)

type Order struct {
	ID       int
	Product  string
	Quantity int
	Status   Status
}

var orderIDCounter = 0
var orderIDMutex sync.Mutex

// NewOrder generates a new order with a unique ID
func NewOrder(product string, quantity int) Order {
	orderIDMutex.Lock()
	orderIDCounter++
	orderID := orderIDCounter
	orderIDMutex.Unlock()
	return Order{ID: orderID, Product: product, Quantity: quantity, Status: Processing}
}

// ProcessOrders processes orders from the queue
func ProcessOrders(orderQueue <-chan Order, inventory *Inventory, wg *sync.WaitGroup) {
	defer wg.Done()
	for order := range orderQueue {
		utils.Log(fmt.Sprintf("Before Processing Order #%d: %s, Stock: %d", order.ID, order.Product, inventory.GetStock(order.Product)))

		if inventory.UpdateStock(order.Product, order.Quantity) {
			time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second) // Simulate processing time
			order.Status = Shipped
			utils.Log(fmt.Sprintf("✅ Order #%d Processed: %s x%d, Remaining Stock: %d", order.ID, order.Product, order.Quantity, inventory.GetStock(order.Product)))
		} else {
			order.Status = Failed
			utils.Log(fmt.Sprintf("❌ Order #%d Failed (Out of Stock): %s x%d, Remaining Stock: %d", order.ID, order.Product, order.Quantity, inventory.GetStock(order.Product)))
		}
	}
}
