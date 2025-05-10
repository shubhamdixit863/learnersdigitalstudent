package model

type Employee struct {
	ID         int      `json:"id"`
	Name       string   `json:"name"`
	Age        int      `json:"age"`
	Department string   `json:"department"`
	Salary     float64  `json:"salary"`
	Address    Address  `json:"address"`
	Skills     []string `json:"skills"`
}
type Address struct {
	Street  string
	City    string
	ZipCode string
}
