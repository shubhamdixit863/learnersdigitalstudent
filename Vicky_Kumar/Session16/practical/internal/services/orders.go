package services

import (
	"fmt"
	"practical/internal/utils"
	"sync"
)

var TotalQuantity = 100

const price = 21

type Order struct {
	OrderId       string
	OrderQuantity int
	OrderAmount   int
	OrderStatus   string
}

func NewOrder(quantity int) *Order {
	return &Order{
		OrderStatus:   "Processing",
		OrderQuantity: quantity,
	}
}
func ProcessOrder(ordersCh chan Order, logCh chan Order, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	for order := range ordersCh {
		order.OrderId = utils.RandomID()
		order.OrderAmount = price * order.OrderQuantity
		mu.Lock()
		if TotalQuantity >= order.OrderQuantity {
			TotalQuantity -= order.OrderQuantity
			order.OrderStatus = "Completed"
		} else {
			order.OrderStatus = "Failed"
		}
		mu.Unlock()
		logCh <- order
	}
	close(logCh)
}

func LogOrder(logCh chan Order, wg *sync.WaitGroup) {
	defer wg.Done()
	for order := range logCh {

		fmt.Printf("Order ID: %s\n", order.OrderId)
		fmt.Printf("Order Quantity: %d\n", order.OrderQuantity)
		fmt.Printf("Order Amount: %d\n", order.OrderAmount)
		fmt.Printf("Order Status: %s\n", order.OrderStatus)
	}
}
