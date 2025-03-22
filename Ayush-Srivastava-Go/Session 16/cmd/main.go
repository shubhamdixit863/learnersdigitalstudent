package main

import (
	"concurrent_order_processing/services"
	"math/rand"
	"time"
)

func main() {
	rand.NewSource(time.Now().UnixNano())
	processor := services.NewOrderProcessor(10) 

	go processor.ProcessOrders()

	for i := 1; i <= 100; i++ {
		go processor.PlaceOrder(i)
		time.Sleep(time.Millisecond * 10) 
	}

	time.Sleep(time.Second * 5) 
	processor.Shutdown()
}
