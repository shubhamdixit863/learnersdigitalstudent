package main

import (
	"concurrent_order_processing_system/internal/models"
	"concurrent_order_processing_system/internal/services"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// rand.Seed(time.Now().UnixNano())
	source := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(source)

	system := services.NewECommerceSystem()

	system.AddProduct(1, "Laptop", 99900.99, 20)
	system.AddProduct(2, "Smartphone", 49900.99, 50)
	system.AddProduct(3, "Headphones", 790.99, 100)
	system.AddProduct(4, "Monitor", 29900.99, 30)
	system.AddProduct(5, "Keyboard", 590.99, 75)

	system.PrintInventory()

	system.StartProcessing()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	for i := 0; i < 30; i++ {
		go func(customerID int) {
			items := []models.OrderItem{
				{ProductID: 1 + rnd.Intn(5), Quantity: 1 + rnd.Intn(3)},
			}

			if rnd.Intn(3) == 0 {
				items = append(items, models.OrderItem{
					ProductID: 1 + rnd.Intn(5),
					Quantity:  1 + rnd.Intn(2),
				})
			}

			orderID, err := system.PlaceOrder(items)
			if err != nil {
				log.Printf("Customer #%d order failed: %v", customerID, err)
			} else {
				log.Printf("Customer #%d placed order #%d", customerID, orderID)
			}
		}(i + 1)

		time.Sleep(time.Duration(rnd.Intn(100)) * time.Millisecond)
	}

	select {
	case <-sigChan:
		log.Println("Shutdown signal received")
	case <-time.After(15 * time.Second):
		log.Println("Simulation complete")
	}

	system.InitiateShutdown()

	system.PrintInventory()
}
