package main

import (
	"math/rand"
	"practical/internal/services"
	"sync"
)

func main() {
	ordersCh := make(chan services.Order)
	logCh := make(chan services.Order)
	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 30; i++ {
			qty := rand.Intn(10) + 1
			order := services.NewOrder(qty)
			ordersCh <- *order
		}
		close(ordersCh)
	}()

	wg.Add(1)
	go services.ProcessOrder(ordersCh, logCh, wg, mu)
	wg.Add(1)
	go services.LogOrder(logCh, wg)
	wg.Wait()
}
