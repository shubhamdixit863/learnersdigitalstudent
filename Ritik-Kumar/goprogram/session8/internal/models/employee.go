package models

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
	Street  string `json:"street"`
	City    string `json:"city"`
	ZipCode string `json:"zipCode"`
}

func (e *Employee) UpdateSalary() {
	if e.Age < 30 {
		e.Salary *= 1.15
	} else if e.Age > 50 {
		e.Salary *= 1.10
	}
}
