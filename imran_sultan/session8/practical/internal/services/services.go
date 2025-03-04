package services

import (
	"encoding/json"
	"os"
	"practical/internal/models"
)

type EmployeeServices struct {
	employee []models.Employee
}

func NewEmployeeServices() *EmployeeServices {
	return &EmployeeServices{}
}

func (e *EmployeeServices) Readjsonfile(path string) (*EmployeeServices, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err

	}
	var emp []models.Employee
	er1 := json.Unmarshal(file, &emp)
	if er1 != nil {
		return nil, er1
	}
	return &EmployeeServices{
		employee: emp,
	}, nil

}
func (es *EmployeeServices) IncreaseSalary() {

	for i := range es.employee {
		if es.employee[i].Age < 30 {
			es.employee[i].Salary += (15.0 / 100.0) * es.employee[i].Salary
		} else if es.employee[i].Age >= 30 && es.employee[i].Age < 50 {
			continue
		} else {
			es.employee[i].Salary += (10.0 / 100.0) * es.employee[i].Salary
		}
	}
}
func (es *EmployeeServices) SortbySalary() {

	for i := range es.employee {
		for j := range es.employee {
			if es.employee[i].Salary > es.employee[j].Salary {
				es.employee[i], es.employee[j] = es.employee[j], es.employee[i]
			}
		}
	}
}
