package services

type Order struct {
	OrderID int
	Items   map[string]int
	Status  string
}

func CreateNewOrder(orderID int, items map[string]int) *Order {
	return &Order{OrderID: orderID, Items: items, Status: "Processing"}
}

func (o *Order) PlaceOrder(orderQueue chan Order) {
	orderQueue <- *o
}
