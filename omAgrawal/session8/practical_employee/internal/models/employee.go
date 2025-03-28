package models

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	ZipCode int    `json:"zip_Code"`
}
type Employee struct {
	Id         int      `json:"id"`
	Name       string   `json:"name"`
	Age        int      `json:"age"`
	Department string   `json:"department"`
	Salary     float64  `json:"salary"`
	Address    Address  `json:"address"`
	Skills     []string `json:"skills"`
}
