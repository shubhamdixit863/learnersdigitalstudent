package inventory

import (
	"sync"
)

var Inventory = map[string]int{}
var InvMutex sync.Mutex

func InitInventory() {
	Inventory["itemA"] = 20
	Inventory["itemB"] = 15
	Inventory["itemC"] = 10
}

func UpdateInventory(item string, quantity int) bool {
	InvMutex.Lock()
	defer InvMutex.Unlock()
	if Inventory[item] >= quantity {
		Inventory[item] -= quantity
		return true
	}
	return false
}
