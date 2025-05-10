package model

type Order struct {
	ID     int
	Status string
	Name   string
}

func NewOrder(id int, status string, Name string) *Order {
	return &Order{
		ID:     id,
		Status: status,
		Name:   Name,
	}
}
