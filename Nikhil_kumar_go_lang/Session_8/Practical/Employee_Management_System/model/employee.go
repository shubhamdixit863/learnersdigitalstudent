package model

type Employee struct {
	Id         int      `json:"id"`
	Name       string   `json:"name"`
	Age        int      `json:"age"`
	Department string   `json:"department"`
	Salary     float64  `json:"salary"`
	Skills     []string `json:"skills"`
	Address    Address  `json:"address"`
}

type Address struct {
	Id      int    `json:"id"`
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zipCode"`
}
