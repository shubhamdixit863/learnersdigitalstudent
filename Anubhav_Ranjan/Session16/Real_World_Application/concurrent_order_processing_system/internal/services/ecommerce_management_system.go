package services

import (
	"concurrent_order_processing_system/internal/models"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
	mu    sync.Mutex
}

type ECommerceSystem struct {
	inventory     map[int]*Product
	orders        map[int]*models.Order
	orderQueue    chan models.Order
	statusUpdates chan models.StatusUpdate
	nextOrderID   int
	orderIDMu     sync.Mutex
	wg            sync.WaitGroup
	shutdown      chan struct{}
}

func NewECommerceSystem() *ECommerceSystem {
	return &ECommerceSystem{
		inventory:     make(map[int]*Product),
		orders:        make(map[int]*models.Order),
		orderQueue:    make(chan models.Order, 10),
		statusUpdates: make(chan models.StatusUpdate, 10),
		nextOrderID:   1,
		shutdown:      make(chan struct{}),
	}
}

func (ecs *ECommerceSystem) AddProduct(id int, name string, price float64, stock int) {
	ecs.inventory[id] = &Product{
		ID:    id,
		Name:  name,
		Price: price,
		Stock: stock,
	}
}

func (ecs *ECommerceSystem) GenerateOrderID() int {
	ecs.orderIDMu.Lock()
	defer ecs.orderIDMu.Unlock()
	id := ecs.nextOrderID
	ecs.nextOrderID++
	return id
}

func (ecs *ECommerceSystem) PlaceOrder(items []models.OrderItem) (int, error) {
	orderID := ecs.GenerateOrderID()

	select {
	case <-ecs.shutdown:
		return 0, fmt.Errorf("system is shutting down, not accepting new orders")
	default:
	}

	for _, item := range items {
		product, exists := ecs.inventory[item.ProductID]
		if !exists {
			return 0, fmt.Errorf("product ID %d does not exist", item.ProductID)
		}

		product.mu.Lock()
		if product.Stock < item.Quantity {
			product.mu.Unlock()
			return 0, fmt.Errorf("insufficient stock for product ID %d", item.ProductID)
		}
		product.mu.Unlock()
	}

	order := models.Order{
		ID:        orderID,
		Items:     items,
		Status:    "Order Placed",
		Timestamp: time.Now(),
	}

	select {
	case ecs.orderQueue <- order:
		log.Printf("Order #%d placed successfully with %d items", orderID, len(items))
		return orderID, nil
	case <-time.After(5 * time.Second):
		return 0, fmt.Errorf("system is busy, please try again later")
	}
}

func (ecs *ECommerceSystem) ProcessOrder(order models.Order) {
	defer ecs.wg.Done()

	log.Printf("Processing order #%d", order.ID)

	ecs.statusUpdates <- models.StatusUpdate{OrderID: order.ID, Status: "Processing"}

	success := true
	reducedStock := make(map[int]int)

	for _, item := range order.Items {
		product := ecs.inventory[item.ProductID]

		product.mu.Lock()
		if product.Stock >= item.Quantity {
			product.Stock -= item.Quantity
			reducedStock[item.ProductID] = item.Quantity
			product.mu.Unlock()
		} else {
			product.mu.Unlock()
			success = false
			break
		}
	}

	if !success {
		for productID, quantity := range reducedStock {
			product := ecs.inventory[productID]
			product.mu.Lock()
			product.Stock += quantity
			product.mu.Unlock()
		}
		log.Printf("Order #%d failed due to inventory changes during processing", order.ID)
		return
	}

	processingTime := time.Duration(1+rand.Intn(3)) * time.Second
	time.Sleep(processingTime)

	ecs.statusUpdates <- models.StatusUpdate{OrderID: order.ID, Status: "Shipped"}

	shippingTime := time.Duration(1+rand.Intn(2)) * time.Second
	time.Sleep(shippingTime)

	ecs.statusUpdates <- models.StatusUpdate{OrderID: order.ID, Status: "Completed"}

	log.Printf("Order #%d completed successfully", order.ID)
}

func (ecs *ECommerceSystem) StatusLoggerUpdater() {
	for update := range ecs.statusUpdates {
		order, exists := ecs.orders[update.OrderID]
		if !exists {
			log.Printf("Order ID %d does not exist", update.OrderID)
		}
		order.Status = update.Status
		log.Printf("Order #%d status: %s", order.ID, order.Status)
	}
}

func (ecs *ECommerceSystem) StartProcessing() {
	go ecs.StatusLoggerUpdater()

	for i := 0; i < 10; i++ {
		go func(workerID int) {
			log.Printf("Order processor worker #%d started", workerID)
			for {
				select {
				case order, ok := <-ecs.orderQueue:
					if !ok {
						log.Printf("Worker #%d shutting down", workerID)
						return
					}
					ecs.wg.Add(1)
					ecs.orders[order.ID] = &order
					ecs.ProcessOrder(order)
				case <-ecs.shutdown:

					// log.Printf("Worker #%d received shutdown signal", workerID)
					// return
					log.Printf("Worker #%d received shutdown signal, finishing remaining orders", workerID)

					for order := range ecs.orderQueue {
						ecs.wg.Add(1)
						ecs.orders[order.ID] = &order
						ecs.ProcessOrder(order)
					}
					log.Printf("Worker #%d finished draining queue after shutdown", workerID)
					return
				}
			}
		}(i + 1)
	}
}

func (ecs *ECommerceSystem) InitiateShutdown() {
	log.Println("Initiating graceful shutdown...")

	close(ecs.shutdown)

	time.Sleep(100 * time.Millisecond)
	close(ecs.orderQueue)

	log.Println("Waiting for all orders to complete...")
	ecs.wg.Wait()

	close(ecs.statusUpdates)

	log.Println("Shutdown complete. All orders processed.")
}

func (ecs *ECommerceSystem) PrintInventory() {
	fmt.Println("\nCurrent Inventory:")
	fmt.Println("------------------")
	for _, product := range ecs.inventory {
		// product.mu.Lock()
		fmt.Printf("Product #%d: %s - Price: Rs.%.2f - Stock: %d\n",
			product.ID, product.Name, product.Price, product.Stock)
		// product.mu.Unlock()
	}
	fmt.Println()
}
