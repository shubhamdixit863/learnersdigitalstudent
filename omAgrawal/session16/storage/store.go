// package storage
//
// import "sync"
//
//	type Order struct {
//		UniqueId int
//		Qunty    int
//		Status   string
//		Price    int
//	}
//
//	func Neworder(UniqueId int, qnty int, status string) *Order {
//		return &Order{UniqueId, qnty, status, Price * qnty}
//	}
//
// var QuantityWarehouse = 10000
// var Price = 100
// var OrderQueue = make(chan *Order)
// var OrderList = make([]*Order, 0)
// var Mu sync.Mutex
// var Wg sync.WaitGroup
package storage

import "sync"

var Mu sync.Mutex
var QuantityWarehouse = 10000
var OrderList = make([]*Order, 0)
var Price = 100

type Order struct {
	UniqueId int
	Qunty    int
	Status   string
	Price    int
}

func Neworder(UniqueId int, qnty int, status string) *Order {
	return &Order{UniqueId, qnty, status, Price * qnty}
}
