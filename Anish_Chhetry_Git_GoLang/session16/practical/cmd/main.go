package main

import (
	"practical/internal/model"
	"practical/internal/services"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	inventory := make(map[string]int)
	inventory["shoe"] = 8
	inventory["hat"] = 23
	inventory["shirt"] = 5

	orderQueue := make(chan model.Order)
	go services.TrackOrder(orderQueue)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		order := model.NewOrder(i+1, "", "shoe")
		go services.ProcessOrder(order, inventory, &mu, &wg, orderQueue)
		time.Sleep(4 * time.Second)
	}
	wg.Wait()
	close(orderQueue)
}
