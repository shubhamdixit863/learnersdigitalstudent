package models

type Employee struct {
	Id         int      `json:"id"`
	Name       string   `json:"name"`
	Age        int      `json:"age"`
	Department string   `json:"department"`
	Salary     float64  `json:"salary"`
	Address    Address  `json:"address"`
	Skillls    []string `json:"skills"`
}
type Address struct {
	Street  string `json: "street"`
	City    string `json: "city"`
	ZipCode string `json:"zipcode"`
}
