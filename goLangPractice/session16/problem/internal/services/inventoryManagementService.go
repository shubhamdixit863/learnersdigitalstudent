package services

import (
	"log"
	"strconv"
	"sync"
)

var stock = map[string]int{
	"Mobile":          5,
	"Laptop":          5,
	"Charger":         5,
	"Cable":           5,
	"Washing Machine": 5,
	"Shirt":           5,
	"Jeans":           5,
}

func UpdateInventory(o Order, orderStatus chan []string, mu *sync.Mutex) {
	//mu.Lock()
	//defer mu.Unlock()

	for item, quantity := range o.Items {
		if availableQuantity, ok := stock[item]; ok {
			if quantity > availableQuantity {
				log.Printf("Item %s is out of stock. Therefore orderId %d is rejected", item, o.OrderID)
				status := []string{"Rejected", strconv.Itoa(o.OrderID)}
				orderStatus <- status
				return
			}
		} else {
			log.Printf("Item %s not found in inventory. Therefore orderId %d is rejected", item, o.OrderID)
			status := []string{"Rejected", strconv.Itoa(o.OrderID)}
			orderStatus <- status
			return
		}
	}

	for item, quantity := range o.Items {
		stock[item] -= quantity
	}

	status := []string{"Shipped", strconv.Itoa(o.OrderID)}
	orderStatus <- status
	status = []string{"Completed", strconv.Itoa(o.OrderID)}
	orderStatus <- status
}
