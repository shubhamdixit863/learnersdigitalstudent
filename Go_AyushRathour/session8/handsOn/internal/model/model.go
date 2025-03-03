package model

type Address struct {
	Street  string `json:"street,omitempty"`
	City    string `json:"city,omitempty"`
	ZipCode string `json:"code,omitempty"`
}

type Employee struct {
	ID         int     `json:"id,omitempty"`
	Name       string  `json:"name,omitempty"`
	Age        int     `json:"age,omitempty"`
	Department string  `json:"dept,omitempty"`
	Salary     float64 `json:"salary,omitempty"`
	Address   Address  `json:"address,omitempty"`
	Skills []string `json:"skills,omitempty"`
}