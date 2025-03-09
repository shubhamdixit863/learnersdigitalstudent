package models

import "encoding/json"

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	ZipCode string `json:"zipCode"`
}

type Employee struct {
	ID         int      `json:"id"`
	Name       string   `json:"name"`
	Age        int      `json:"age"`
	Department string   `json:"department"`
	Salary     float64  `json:"salary"`
	Address    Address  `json:"address"`
	Skills     []string `json:"skills"`
}

func (e Employee) String() string {
	data, _ := json.MarshalIndent(e, "", "  ")
	return string(data)
}
