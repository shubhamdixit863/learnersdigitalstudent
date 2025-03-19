package order

type Order struct {
	OrderID  int
	Item     string
	Quantity int
	Status   string
}

var orderIDCount int
var OrderQueue = make(chan Order, 100)

func generateOrderID() int {
	orderIDCount++
	return orderIDCount
}

func PlaceOrder(item string, quantity int) {
	order := Order{
		OrderID:  generateOrderID(),
		Item:     item,
		Quantity: quantity,
		Status:   "Processing",
	}
	OrderQueue <- order
}
