package services

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type OrderProcessor struct {
	orderQueue chan Order
	inventory  *Inventory
	wg         sync.WaitGroup
	shutdown   chan struct{}
}

func NewOrderProcessor(stock int) *OrderProcessor {
	return &OrderProcessor{
		orderQueue: make(chan Order, 100),
		inventory:  NewInventory(stock),
		shutdown:   make(chan struct{}),
	}
}

func (op *OrderProcessor) PlaceOrder(orderID int) {
	select {
	case <-op.shutdown:
		fmt.Printf("Order %d rejected: System shutting down\n", orderID)
	default:
		op.orderQueue <- Order{ID: orderID, Status: "Processing"}
		fmt.Printf("Order %d placed.\n", orderID)
	}
}

func (op *OrderProcessor) ProcessOrders() {
	for order := range op.orderQueue {
		op.wg.Add(1)
		go op.processOrder(order)
	}
}

func (op *OrderProcessor) processOrder(order Order) {
	defer op.wg.Done()

	if op.inventory.DeductStock() {
		fmt.Printf("Order %d is being processed. Remaining stock: %d\n", order.ID, op.inventory.GetStock())
		time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second) 
		fmt.Printf("Order %d shipped\n", order.ID)
	} else {
		fmt.Printf("Order %d rejected: Out of stock\n", order.ID)
	}
}

func (op *OrderProcessor) Shutdown() {
	close(op.shutdown) 
	close(op.orderQueue)
	op.wg.Wait() 
	fmt.Println("All orders processed. System shutting down.")
}
