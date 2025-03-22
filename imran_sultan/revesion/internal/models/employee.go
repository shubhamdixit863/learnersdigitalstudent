package models

type Employee struct {
	Id         int     `json:"Id"`
	Name       string  `json:"name"`
	Age        int     `json:"Age"`
	Department string  `json:"department"`
	Salary     float64 `json:"Salary"`
	Address    Address `json:"address"`
}
type Address struct {
	Street  string   `json:"street"`
	City    string   `json:"city"`
	Zipcode string   `json:"zipcode"`
	Skills  []string `json:"skills"`
}
